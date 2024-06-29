package models

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
