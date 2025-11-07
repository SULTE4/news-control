// service.go
package main

import (
	"fmt"

	"github.com/sulte4/news-control/observer"
	"github.com/sulte4/news-control/strategy"
)

type ArticleService struct {
	agency *observer.NewAgency
	ws     *WSManager
}

func NewArticleService(agency *observer.NewAgency, ws *WSManager) *ArticleService {
	return &ArticleService{agency: agency, ws: ws}
}

func (s *ArticleService) CreateArticle(title, content string) {
	article := observer.Article{
		Title:   title,
		Content: content,
		InStock: true,
	}

	fmt.Println("[SERVICE] Created article:", title)
	s.agency.UpdateAvailability(article)

	s.ws.Broadcast(article)
}

func (s *ArticleService) RegisterUser(name, email, method string) *observer.User {
	user := observer.NewUser(name)
	user.SetEmail(email)
	user.SetNotifier(strategy.NotificationFactory(method))
	s.agency.Register(user)
	return user
}

func (s *ArticleService) DeregisterUser(name string) {
	for _, u := range s.agency.Subscribers() {
		if u.GetName() == name {
			s.agency.Deregister(u)
			return
		}
	}
}
