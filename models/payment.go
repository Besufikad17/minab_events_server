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
