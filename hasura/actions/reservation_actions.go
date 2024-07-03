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
		Query: constants.ReserveEvent,
		Variables: models.ReserveEventArgs{
			Event_id:  variables.Event_id,
			User_id:   variables.User_id,
			Status:    variables.Status,
			Ticket_id: variables.Ticket_id,
		},
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
