package main

import (
	"log"
	"net/http"

	handlers "github.com/Besufikad17/minab_events/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/Register", handlers.RegisterHandler)

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
