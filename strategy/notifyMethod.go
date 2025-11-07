package strategy

type NotifyMethod interface {
	Notify(userEmail, content string)
}
