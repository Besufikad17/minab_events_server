package hasura

type GraphQLError struct {
	Message string `json:"message"`
}

type ActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            RegisterArgs           `json:"input"`
}

type LoginInput struct {
	LoginText string `json:"login_text"`
	Password  string `json:"password"`
}

type NewUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
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

type RegisterOutput struct {
	Email        *string `json:"email"`
	First_name   *string `json:"first_name"`
	Id           int     `json:"id"`
	Last_name    *string `json:"last_name"`
	Phone_number *string `json:"phone_number"`
}

type RegisterArgs struct {
	First_name   string
	Last_name    string
	Email        string
	Phone_number string
	Password     string
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type RegisterGraphQLData struct {
	Insert_users_one RegisterOutput `json:"insert_users_one"`
}

type RegisterGraphQLResponse struct {
	Data   RegisterGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError      `json:"errors,omitempty"`
}
