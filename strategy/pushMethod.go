package strategy

import "fmt"

type PushNotification struct{}

func (p *PushNotification) Notify(userEmail, content string) {
	fmt.Printf("[PUSH] Notification to %s: %s\n", userEmail, content)
}
