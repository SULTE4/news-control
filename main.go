package main

import (
	"github.com/sulte4/news-control/facade"
	"github.com/sulte4/news-control/observer"
	"github.com/sulte4/news-control/strategy"
)

func main() {
	agency := observer.GetInstance()

	// Create users
	john := observer.NewUser("John")
	john.SetEmail("john@example.com")
	john.SetNotifier(strategy.NotificationFactory("email"))

	anna := observer.NewUser("Anna")
	anna.SetEmail("anna@example.com")
	anna.SetNotifier(strategy.NotificationFactory("push"))

	agency.Register(john)
	agency.Register(anna)

	// Use Facade to publish
	f := facade.NewAddingArticleFacade(agency)
	f.AddArticle("Breaking News", "New design patterns in Go!")

	// Anna switches strategy
	anna.SetNotifier(strategy.NotificationFactory("email"))
	f.AddArticle("Tech Update", "Strategy pattern explained.")
}
