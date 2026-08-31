package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/Besufikad17/minab_events/handlers"
	"github.com/Besufikad17/minab_events/hasura/events"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Printf("Error loading .env file: %v", err)
	}

	mux := http.NewServeMux()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// auth handlers
	mux.HandleFunc("/Register", handlers.RegisterHandler)
	mux.HandleFunc("/Login", handlers.LoginHandler)

	// event handlers
	mux.HandleFunc("/events/Create", handlers.CreateEventHandler)
	mux.HandleFunc("/events/Reserve", handlers.ReserveEvent)
	mux.HandleFunc("/events/Reserved", events.NotifyUser)

	// image handlers
	mux.HandleFunc("/image/Add", handlers.AddImagesHandler)

	// file serving
	mux.HandleFunc("/images/", handlers.ServeImage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	if port[0] != ':' {
		port = ":" + port
	}

	log.Fatal(http.ListenAndServe(port, mux))
}
