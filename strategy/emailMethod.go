package strategy

import "fmt"

type EmailNotification struct{}

func (e *EmailNotification) Notify(userEmail, content string) {
	fmt.Printf("[EMAIL] Sent to %s: %s\n", userEmail, content)
}
