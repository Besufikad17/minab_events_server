package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	handlers "github.com/Besufikad17/minab_events/handlers"
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

	mux.HandleFunc("/Register", handlers.RegisterHandler)

	err = http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
