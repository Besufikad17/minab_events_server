package actions

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	models "github.com/Besufikad17/minab_events/models"
	constants "github.com/Besufikad17/minab_events/utils/constants"
)

func CreateEvent(args map[string]interface{}) (response models.CreateEventOutput, err error) {
	hasuraResponse, err := executeCreateEvent(args)
	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Insert_events_one
	return
}
func executeCreateEvent(variables map[string]interface{}) (response *models.CreateEventGraphQLResponse, err error) {

	reqBody := models.GraphQLRequest{
		Query:     constants.CreateEvent,
		Variables: variables,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	hasuraURL := os.Getenv("HASURA_URL")
	resp, err := http.Post(hasuraURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, err
	}

	return
}
