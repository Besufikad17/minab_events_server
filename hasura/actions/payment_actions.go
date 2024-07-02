package actions

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Besufikad17/minab_events/models"
)

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
