package models

type CreateLocationActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            CreateLocationArgs     `json:"input"`
}

type CreateLocationArgs struct {
	City  string
	Venue string
	Lat   float32
	Lng   float32
}

type CreateLocationGraphQLData struct {
	Insert_locations_one CreateLocationOutput `json:"insert_locations_one"`
}

type CreateLocationGraphQLResponse struct {
	Data   CreateLocationGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError            `json:"errors,omitempty"`
}

type CreateLocationOutput struct {
	City  string
	Id    int
	Venue string
}
