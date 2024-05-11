package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	actions "github.com/Besufikad17/minab_events/hasura/actions"
	models "github.com/Besufikad17/minab_events/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	var actionPayload models.ActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	result, err := actions.Register(actionPayload.Input)

	if err != nil {
		errorObject := models.GraphQLError{
			Message: err.Error(),
		}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}

	data, _ := json.Marshal(result)
	w.Write(data)
}
