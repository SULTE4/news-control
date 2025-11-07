package observer

type Observer interface {
	Update(item string)
	GetName() string
	GetEmail() string
}
