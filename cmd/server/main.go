package main

import (
	"log"
	"net/http"

	"github.com/amjadfqs/go-htmx-starter/internal/handlers"
)

func main() {
	// Static file server
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Route handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/api/example", handlers.ExampleHtmxHandler)
	http.HandleFunc("/api/client-info", handlers.ClientInfoHandler)

	// Start the server
	log.Println("Server starting at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
