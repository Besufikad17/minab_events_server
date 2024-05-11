package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	handlers "github.com/Besufikad17/minab_events/handlers"
)

func main() {
	mux := http.NewServeMux()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	mux.HandleFunc("/Register", handlers.RegisterHandler)

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
