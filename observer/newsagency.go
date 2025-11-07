package observer

import (
	"fmt"
	"sync"
)

type NewAgency struct {
	subcribers []Observer
}

var (
	instance *NewAgency
	once     sync.Once
)

func GetInstance() *NewAgency {
	once.Do(func() {
		instance = &NewAgency{subcribers: []Observer{}}
	})
	return instance
}

func (a *NewAgency) Register(subscriber Observer) {
	a.subcribers = append(a.subcribers, subscriber)
	fmt.Println(subscriber.GetName(), " registered")
}

func (a *NewAgency) Deregister(subscriber Observer) {
	a.deleteFromList(subscriber)
}

func (a *NewAgency) UpdateAvailability(article Article) {
	a.NotifyAll(article)
}

func (a *NewAgency) NotifyAll(article Article) {
	for _, subscriber := range a.subcribers {
		subscriber.Update(article.Title)
	}
}

func (a *NewAgency) deleteFromList(subscriber Observer) {
	for i, sub := range a.subcribers {
		if sub.GetName() == subscriber.GetName() {
			a.subcribers = append(a.subcribers[:i], a.subcribers[i+1:]...)
			fmt.Println(subscriber.GetName(), " unregistered")
		}
	}
}
