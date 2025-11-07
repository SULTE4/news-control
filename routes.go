package main

import (
	"encoding/json"
	"net/http"
)

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type RegisterRequest struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Method string `json:"method"`
}

func SetupRoutes(wsManager *WSManager, service *ArticleService) {
	http.HandleFunc("/ws", wsManager.HandleConnection)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user := service.RegisterUser(req.Name, req.Email, req.Method)
		w.Write([]byte("Registered " + user.GetName()))
	})

	http.HandleFunc("/deregister", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			http.Error(w, "missing ?name=", http.StatusBadRequest)
			return
		}
		service.DeregisterUser(name)
		w.Write([]byte("Unregistered " + name))
	})

	http.HandleFunc("/article", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req CreateArticleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		service.CreateArticle(req.Title, req.Content)
		w.Write([]byte("Article created"))
	})
}
