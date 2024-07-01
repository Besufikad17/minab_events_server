package actions

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/Besufikad17/minab_events/models"
	"github.com/Besufikad17/minab_events/utils/constants"
)

func AddTicket(args models.AddTicketArgs, token string) (response models.AddTicketOutput, err error) {
	hasuraResponse, err := executeAddTicket(args, token)

	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Insert_tickets_one
	return
}

func executeAddTicket(variables models.AddTicketArgs, token string) (response *models.AddTicketGraphQLResponse, err error) {
	reqBody := models.AddTicketGraphQLRequest{
		Query:     constants.AddTicket,
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
