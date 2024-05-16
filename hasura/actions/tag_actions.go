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

func CreateTag(args models.CreateTagArgs) (response models.CreateTagOutput, err error) {
	hasuraResponse, err := executeCreateTag(args)

	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Insert_tags_one
	return
}

func executeCreateTag(variables models.CreateTagArgs) (response *models.CreateTagGraphQLResponse, err error) {
	reqBody := models.GraphQLRequest{
		Query: constants.CreateTag,
		Variables: map[string]interface{}{
			"name":     variables.Name,
			"event_id": variables.Event_id,
		},
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	hasuraUrl := os.Getenv("HASURA_URL")
	resp, err := http.Post(hasuraUrl, "application/json", bytes.NewBuffer(reqBytes))
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
