package models

type LoginActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            LoginArgs              `json:"input"`
}

type LoginArgs struct {
	Login_text  string
	Password    string
	Remember_me bool
}

type LoginInput struct {
	LoginText string `json:"login_text"`
	Password  string `json:"password"`
}

type LoginOutput struct {
	Token string
}

type RegisterActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            RegisterArgs           `json:"input"`
}

type RegisterArgs struct {
	First_name   string
	Last_name    string
	Email        string
	Phone_number string
	Password     string
	Remember_me  bool
}

type RegisterGraphQLData struct {
	Insert_users_one RegisterOutput `json:"insert_users_one"`
}

type RegisterGraphQLResponse struct {
	Data   RegisterGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError      `json:"errors,omitempty"`
}

type RegisterOutput struct {
	Email        *string `json:"email"`
	First_name   *string `json:"first_name"`
	Id           int     `json:"id"`
	Last_name    *string `json:"last_name"`
	Phone_number *string `json:"phone_number"`
}
