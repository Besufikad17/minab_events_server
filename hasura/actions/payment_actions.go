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

func CreatePayment(args models.CreatePaymentArgs, token string) (response *models.CreatePaymentOutput, err error) {
	hasuraResponse, err := executeCreatePayment(args, token)

	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = &hasuraResponse.Data.Insert_payments_one
	return
}

func executeCreatePayment(variables models.CreatePaymentArgs, token string) (response *models.CreatePaymentGraphQLResponse, err error) {
	reqBody := models.CreatePaymentGraphQLRequest{
		Query:     constants.CreatePayment,
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

func PayWithChapa(args models.ChapaArgs) (response *models.ChapaOutput, err error) {
	reqBytes, err := json.Marshal(args)
	if err != nil {
		return
	}

	chapaURL := os.Getenv("CHAPA_URL")
	chapaAuth := os.Getenv("CHAPA_AUTH")

	client := &http.Client{}
	req, err := http.NewRequest("POST", chapaURL, bytes.NewBuffer(reqBytes))
	req.Header.Add("Authorization", "Bearer "+chapaAuth)
	req.Header.Add("Content-Type", "application/json")

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
