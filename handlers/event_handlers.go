package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	actions "github.com/Besufikad17/minab_events/hasura/actions"
	models "github.com/Besufikad17/minab_events/models"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var actionPayload models.CreateEventActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	location, err := actions.CreateLocation(models.CreateLocationArgs{
		City:  actionPayload.Input.City,
		Venue: actionPayload.Input.Venue,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mapVariables := map[string]interface{}{
		"title":         actionPayload.Input.Title,
		"description":   actionPayload.Input.Description,
		"image":         actionPayload.Input.Image,
		"user_id":       actionPayload.Input.User_id,
		"category_id":   actionPayload.Input.Category_id,
		"location_id":   location.Id,
		"start_date":    actionPayload.Input.Start_date,
		"end_date":      actionPayload.Input.End_date,
		"enterance_fee": actionPayload.Input.Enterance_fee,
	}

	result, err := actions.CreateEvent(mapVariables)
	fmt.Println(actionPayload.Input.Tags)

	if err != nil {
		errorObject := models.GraphQLError{
			Message: err.Error(),
		}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}

	for _, tag := range actionPayload.Input.Tags {
		_, err = actions.CreateTag(models.CreateTagArgs{
			Name:     tag,
			Event_id: result.Id,
		})

		if err != nil {
			errorObject := models.GraphQLError{
				Message: err.Error(),
			}
			errorBody, _ := json.Marshal(errorObject)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorBody)
			return
		}
	}

	data, _ := json.Marshal(result)
	w.Write(data)
}
