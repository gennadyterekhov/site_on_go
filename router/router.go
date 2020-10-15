package router

import (
	"net/http"
	"portfolio/controllers"
)

func Router() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/reactPage", controllers.ReactPage)
	// http.HandleFunc("/webhook", webhook.Webhook)
}
