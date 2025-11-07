package main

import (
	"fmt"
	"net/http"

	"github.com/sulte4/news-control/observer"
)

func main() {
	agency := observer.GetInstance()
	wsManager := NewWSManager(agency)
	service := NewArticleService(agency, wsManager)

	SetupRoutes(wsManager, service)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
