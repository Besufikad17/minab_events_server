package models

type ChapaArgs struct {
	Amount       float32 `json:"amount"`
	Currency     string  `json:"currency"`
	Email        string  `json:"email"`
	First_name   string  `json:"first_name"`
	Last_name    string  `json:"last_name"`
	Phone_number string  `json:"phone_number"`
	Tx_ref       string  `json:"tx_ref"`
	Callback_url string  `json:"callback_url"`
	Return_url   string  `json:"return_url"`
}

type ChapaData struct {
	Checkout_url string `json:"checkout_url"`
}

type ChapaOutput struct {
	Message string    `json:"message"`
	Status  string    `json:"status"`
	Data    ChapaData `json:"data"`
}

type CreatePaymentActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            CreatePaymentArgs      `json:"input"`
}

type CreatePaymentArgs struct {
	User_id   int     `json:"user_id"`
	Ticket_id int     `json:"ticket_id"`
	Amount    float32 `json:"amount"`
	Status    string  `json:"status"`
	Reference string  `json:"reference"`
}

type CreatePaymentGraphQLRequest struct {
	Query     string            `json:"query"`
	Variables CreatePaymentArgs `json:"variables"`
}

type CreatePaymentGraphQLData struct {
	Insert_payments_one CreatePaymentOutput `json:"insert_payments_one"`
}

type CreatePaymentGraphQLResponse struct {
	Data   CreatePaymentGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError           `json:"errors,omitempty"`
}

type CreatePaymentOutput struct {
	Id int `json:"id"`
}
