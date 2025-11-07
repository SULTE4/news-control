package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/sulte4/news-control/observer"
	"github.com/sulte4/news-control/strategy"
)

type WSManager struct {
	clients map[string]*websocket.Conn
	mu      sync.RWMutex
	agency  *observer.NewAgency
}

func NewWSManager(agency *observer.NewAgency) *WSManager {
	return &WSManager{
		clients: make(map[string]*websocket.Conn),
		agency:  agency,
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// HandleConnection upgrades to WS only if user exists in the agency
func (w *WSManager) HandleConnection(wr http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(wr, "missing ?name=", http.StatusBadRequest)
		return
	}

	// Check user exists
	if !w.agency.Exists(name) {
		fmt.Println("[ ws ] user not found")
		http.Error(wr, "user not registered in agency", http.StatusForbidden)
		return
	}

	conn, err := upgrader.Upgrade(wr, r, nil)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	w.mu.Lock()
	w.clients[name] = conn
	w.mu.Unlock()

	fmt.Printf("[WS] %s connected\n", name)
	conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Welcome %s! Connected to live news feed.", name)))
}

func (w *WSManager) SendMessage(name string, msg string) {
	w.mu.RLock()
	conn, ok := w.clients[name]
	w.mu.RUnlock()
	if !ok {
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (w *WSManager) Broadcast(article observer.Article) {
	w.mu.RLock()
	defer w.mu.RUnlock()

	for name, conn := range w.clients {
		// Find the user's notification method
		var method string
		for _, sub := range w.agency.Subscribers() {
			if sub.GetName() == name {
				// The Observer is a *User, so type assert:
				if u, ok := sub.(*observer.User); ok && u.Notifier != nil {
					switch u.Notifier.(type) {
					case *strategy.EmailNotification:
						method = "email"
					case *strategy.PushNotification:
						method = "push"
					default:
						method = "unknown"
					}
				}
				break
			}
		}

		msg := fmt.Sprintf("📰 New article: %s\nDelivery method: %s", article.Title, method)
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
		fmt.Printf("[WS] Sent to %s (%s)\n", name, method)
	}
}
