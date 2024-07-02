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

func CreateEvent(args map[string]interface{}, token string) (response models.CreateEventOutput, err error) {
	hasuraResponse, err := executeCreateEvent(args, token)
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
func executeCreateEvent(variables map[string]interface{}, token string) (response *models.CreateEventGraphQLResponse, err error) {
	reqBody := models.GraphQLRequest{
		Query:     constants.CreateEvent,
		Variables: variables,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	hasuraURL := os.Getenv("HASURA_URL")
	client := &http.Client{}
	req, err := http.NewRequest("POST", hasuraURL, bytes.NewBuffer(reqBytes))
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
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

func ReserveEvent(args models.ReserveEventArgs, token string) (response *models.ReserveEventOutput, err error) {
	hasuraResponse, err := executeReserveEvent(args, token)

	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = &hasuraResponse.Data.Insert_reservations_one
	return
}

func executeReserveEvent(variables models.ReserveEventArgs, token string) (response *models.ReserveEventGraphQLResponse, err error) {
	reqBody := models.ReserveEventGraphQLRequest{
		Query:     constants.ReserveEvent,
		Variables: variables,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	hasuraURL := os.Getenv("HASURA_URL")
	client := &http.Client{}
	req, err := http.NewRequest("POST", hasuraURL, bytes.NewBuffer(reqBytes))
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return
	}

	return
}
