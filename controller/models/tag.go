package models

type CreateTagActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            CreateTagArgs          `json:"input"`
}

type CreateTagArgs struct {
	Event_id int
	Name     string
}

type CreateTagGraphQLData struct {
	Insert_tags_one CreateTagOutput `json:"insert_tags_one"`
}

type CreateTagGraphQLResponse struct {
	Data   CreateTagGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError       `json:"errors,omitempty"`
}

type CreateTagOutput struct {
	Id   int
	Name string
}
