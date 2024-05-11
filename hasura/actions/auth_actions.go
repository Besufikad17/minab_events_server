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
	helpers "github.com/Besufikad17/minab_events/utils/helpers"
)

func Register(args models.RegisterArgs) (response models.RegisterOutput, err error) {
	hasuraResponse, err := execute(args)
	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Insert_users_one
	return
}

func execute(variables models.RegisterArgs) (response models.RegisterGraphQLResponse, err error) {
	hashedPassword, err := helpers.Hash(variables.Password)

	if err != nil {
		return
	}

	mapVariables := map[string]interface{}{
		"first_name":   variables.First_name,
		"last_name":    variables.Last_name,
		"email":        variables.Email,
		"phone_number": variables.Phone_number,
		"password":     hashedPassword,
	}

	reqBody := models.GraphQLRequest{
		Query:     constants.Register,
		Variables: mapVariables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	hasuraURL := os.Getenv("HASURA_URL")
	resp, err := http.Post(hasuraURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
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
