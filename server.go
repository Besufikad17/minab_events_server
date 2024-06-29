package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/Besufikad17/minab_events/handlers"
	"github.com/Besufikad17/minab_events/hasura/events"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
	mux.HandleFunc("/events/Reserve", events.NotifyUser)

	// image handlers
	mux.HandleFunc("/image/Add", handlers.AddImagesHandler)

	// file serving
	mux.HandleFunc("/images/", handlers.ServeImage)

	err = http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
