package models

type GraphQLError struct {
	Message string `json:"message"`
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}
