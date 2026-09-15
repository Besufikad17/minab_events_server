package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Besufikad17/minab_events/internal"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	Ctx           context.Context
	GraphQLClient internal.GraphQLClient
}

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]any{
		"message": "hi :)",
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Write(data)
}

func NewHandler(ctx context.Context, client internal.GraphQLClient) *Handler {
	return &Handler{
		Ctx:           ctx,
		GraphQLClient: client,
	}
}
