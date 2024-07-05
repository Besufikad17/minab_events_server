package actions

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	models "github.com/Besufikad17/minab_events/models"
	constants "github.com/Besufikad17/minab_events/utils/constants"
	helpers "github.com/Besufikad17/minab_events/utils/helpers"
)

func AddImages(args models.AddImagesArgs, token string) (response *models.AddImagesOutput, err error) {
	var images []models.AddImagesImagesInsertInput

	for _, image := range args.Images {
		if strings.HasPrefix(*image.Url, "http") {
			images = append(images, image)
		} else {
			fileName, err := helpers.SaveImageToFile(*image.Url)
			filePath := "http://localhost:5000/images/" + *fileName
			images = append(images, models.AddImagesImagesInsertInput{
				Event_id: image.Event_id,
				Url:      &filePath,
			})

			if err != nil {
				return nil, err
			}
		}
	}

	args.Images = images

	hasuraResponse, err := executeAddImages(args, token)

	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = &hasuraResponse.Data.Insert_images
	return

}

func executeAddImages(variables models.AddImagesArgs, token string) (response *models.AddImagesGraphQLResponse, err error) {
	reqBody := models.AddImagesGraphQLRequest{
		Query:     constants.AddImages,
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

func CreateImage(args models.CreateImageArgs, token string) (response models.CreateImageOutput, err error) {
	hasuraResponse, err := executeCreateImage(args, token)

	if err != nil {
		return
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}

	response = hasuraResponse.Data.Insert_images_one
	return
}

func executeCreateImage(variables models.CreateImageArgs, token string) (response *models.CreateImageGraphQLResponse, err error) {
	reqBody := models.CreateImageGraphQLRequest{
		Query:     constants.CreateImage,
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
		return nil, err
	}
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, err
	}

	return
}
