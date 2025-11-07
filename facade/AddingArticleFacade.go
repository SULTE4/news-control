package facade

import (
	"github.com/sulte4/news-control/observer"
)

type AddingArticleFacade struct {
	Agency *observer.NewAgency
}

func NewAddingArticleFacade(agency *observer.NewAgency) *AddingArticleFacade {
	return &AddingArticleFacade{Agency: agency}
}

func (f *AddingArticleFacade) AddArticle(title, content string) {
	article := observer.NewArticle(title, content)
	f.Agency.UpdateAvailability(article)
}
