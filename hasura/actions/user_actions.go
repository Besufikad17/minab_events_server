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

func SearchUser(args models.SearchUserArgs) (response models.SearchUserOutput, err error) {
	hasuraResponse, err := execute(args)
	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	if len(hasuraResponse.Data.Users) == 0 {
		err = errors.New("user not found")
		return
	}

	response = hasuraResponse.Data.Users[0]
	return
}

func execute(variables models.SearchUserArgs) (response *models.SearchUserGraphQLResponse, err error) {
	mapVariables := map[string]interface{}{
		"login_text": variables.Login_text,
	}
	reqBody := models.GraphQLRequest{
		Query:     constants.SearchUser,
		Variables: mapVariables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
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

func GetUserById(args map[string]interface{}) (response models.GetUserByIdOutput, err error) {

	hasuraResponse, err := executeGetUserById(args)

	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Users[0]
	return

}

func executeGetUserById(variables map[string]interface{}) (response models.GetUserByIdGraphQLResponse, err error) {
	reqBody := models.GraphQLRequest{
		Query:     constants.GetUserById,
		Variables: variables,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	hasuraUrl := os.Getenv("HASURA_URL")
	resp, err := http.Post(hasuraUrl, "application/json", bytes.NewBuffer(reqBytes))
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
