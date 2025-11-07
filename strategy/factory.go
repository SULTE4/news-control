package strategy

import "fmt"

func NotificationFactory(method string) NotifyMethod {
	switch method {
	case "email":
		return &EmailNotification{}
	case "push":
		return &PushNotification{}
	default:
		fmt.Println("Invalid method")
		return nil
	}
}
