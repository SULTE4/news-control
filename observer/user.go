package observer

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sulte4/news-control/strategy"
)

type User struct {
	id       uuid.UUID
	name     string
	email    string
	Notifier strategy.NotifyMethod
}

func NewUser(name string) *User {
	id := uuid.New()
	return &User{
		id:   id,
		name: name,
	}
}

func (s *User) GetName() string {
	return s.name
}

func (s *User) SetEmail(email string) {
	s.email = email
}

func (s *User) GetEmail() string {
	return s.email
}

func (u *User) SetNotifier(notifier strategy.NotifyMethod) {
	u.Notifier = notifier
}

func (u *User) Update(item string) {
	if u.Notifier == nil {
		fmt.Printf("User %s has no notifier configured!\n", u.name)
		return
	}
	u.Notifier.Notify(u.email, fmt.Sprintf("New article: %s", item))
}
