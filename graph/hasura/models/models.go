package hasura

type GraphQLError struct {
	Message string `json:"message"`
}

type RegisterOutput struct {
	Email        *string
	First_name   *string
	Id           int
	Last_name    *string
	Phone_number *string
}

type Mutation struct {
	Register *RegisterOutput
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

type GraphQLData struct {
	Insert_users_one RegisterOutput `json:"insert_users_one"`
}

type GraphQLResponse struct {
	Data   GraphQLData    `json:"data,omitempty"`
	Errors []GraphQLError `json:"errors,omitempty"`
}
