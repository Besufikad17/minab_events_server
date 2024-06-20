package models

type GetUserByIdActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            GetUserByIdArgs        `json:"input"`
}

type GetUserByIdArgs struct {
	Id int
}

type GetUserByIdGraphQLData struct {
	Users []GetUserByIdOutput `json:"users"`
}

type GetUserByIdOutput struct {
	Email      string
	First_name string
	Last_name  string
}

type GetUserByIdGraphQLResponse struct {
	Data   GetUserByIdGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError         `json:"errors,omitempty"`
}

type NewUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type SearchUserActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            SearchUserArgs         `json:"input"`
}

type SearchUserArgs struct {
	Login_text string
}

type SearchUserGraphQLData struct {
	Users []SearchUserOutput `json:"users"`
}

type SearchUserGraphQLResponse struct {
	Data   SearchUserGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError        `json:"errors,omitempty"`
}

type SearchUserOutput struct {
	Email        string `json:"email"`
	First_name   string `json:"first_name"`
	Id           int    `json:"id"`
	Last_name    string `json:"last_name"`
	Password     string `json:"password"`
	Phone_number string `json:"phone_number"`
}

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Token       string `json:"token"`
}
