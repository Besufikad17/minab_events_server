package models

import (
	"time"
)

type AddImagesActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            AddImagesArgs          `json:"input"`
}

type AddImagesArgs struct {
	Images []AddImagesImagesInsertInput `json:"images"`
}

type AddImagesImagesInsertInput struct {
	Event_id *int    `json:"event_id"`
	Url      *string `json:"url"`
}

type AddImagesGraphQLRequest struct {
	Query     string        `json:"query"`
	Variables AddImagesArgs `json:"variables"`
}

type AddImagesGraphQLData struct {
	Insert_images AddImagesOutput `json:"insert_images"`
}

type AddImagesGraphQLResponse struct {
	Data   AddImagesGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError       `json:"errors,omitempty"`
}

type AddImagesOutput struct {
	Affected_rows int `json:"affected_rows"`
}

type CreateImageActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            CreateImageArgs        `json:"input"`
}

type CreateImageArgs struct {
	Event_id int    `json:"event_id"`
	Url      string `json:"url"`
}

type CreateImageGraphQLRequest struct {
	Query     string          `json:"query"`
	Variables CreateImageArgs `json:"variables"`
}

type CreateImageGraphQLData struct {
	Insert_images_one CreateImageOutput `json:"insert_images_one"`
}

type CreateImageGraphQLResponse struct {
	Data   CreateImageGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError         `json:"errors,omitempty"`
}

type CreateImageOutput struct {
	Id  int
	Url string
}

type ImageAsset struct {
	AssetID          string    `json:"asset_id"`
	PublicID         string    `json:"public_id"`
	Version          int       `json:"version"`
	VersionID        string    `json:"version_id"`
	Signature        string    `json:"signature"`
	Width            int       `json:"width"`
	Height           int       `json:"height"`
	Format           string    `json:"format"`
	ResourceType     string    `json:"resource_type"`
	CreatedAt        time.Time `json:"created_at"`
	Tags             []string  `json:"tags"`
	Bytes            int       `json:"bytes"`
	Type             string    `json:"type"`
	ETag             string    `json:"etag"`
	Placeholder      bool      `json:"placeholder"`
	URL              string    `json:"url"`
	SecureURL        string    `json:"secure_url"`
	Folder           string    `json:"folder"`
	AccessMode       string    `json:"access_mode"`
	OriginalFilename string    `json:"original_filename"`
}
