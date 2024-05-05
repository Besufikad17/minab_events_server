package actions

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	models "github.com/Besufikad17/minab_events/graph/hasura/models"
	constants "github.com/Besufikad17/minab_events/graph/utils/constants"
)

func SearchUser(args models.SearchUserArgs) (response models.SearchUserOutput, err error) {

	hasuraResponse, err := executeSearchUser(args)

	// throw if any unexpected error happens
	if err != nil {
		return
	}

	// delegate Hasura error
	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Search_user[0]
	return

}
func executeSearchUser(variables models.SearchUserArgs) (response models.SearchUserGraphQLResponse, err error) {
	mapVariables := map[string]interface{}{
		"loginText": variables.LoginText,
	}

	// build the request body
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
		return
	}

	// parse the response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return
	}

	// return the response
	return
}
