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

func CreateLocation(args models.CreateLocationArgs) (response models.CreateLocationOutput, err error) {
	hasuraResponse, err := executeCreateLocation(args)

	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Insert_locations_one
	return
}

func executeCreateLocation(variables models.CreateLocationArgs) (response *models.CreateLocationGraphQLResponse, err error) {
	reqBody := models.GraphQLRequest{
		Query: constants.CreateLocation,
		Variables: map[string]interface{}{
			"city":  variables.City,
			"venue": variables.Venue,
		},
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
