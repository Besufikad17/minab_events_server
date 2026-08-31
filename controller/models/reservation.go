package models

type ReserveEventActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            ReserveEventArgs       `json:"input"`
}

type ReserveEventArgs struct {
	User_id   int    `json:"user_id"`
	Event_id  int    `json:"event_id"`
	Ticket_id int    `json:"ticket_id"`
	Status    string `json:"status"`
}

type ReserveEventGraphQLRequest struct {
	Query     string           `json:"query"`
	Variables ReserveEventArgs `json:"variables"`
}

type ReserveEventGraphQLData struct {
	Insert_reservations_one ReserveEventOutput `json:"insert_reservations_one"`
}

type ReserveEventGraphQLResponse struct {
	Data   ReserveEventGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError          `json:"errors,omitempty"`
}

type ReserveEventOutput struct {
	Id          int    `json:"id"`
	CheckoutUrl string `json:"checkoutUrl"`
}

type OnEventReservedActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            OnEventReservedArgs    `json:"input"`
}

type OnEventReservedArgs struct {
	Id        int    `json:"id"`
	User_id   int    `json:"user_id"`
	Event_id  int    `json:"event_id"`
	Ticket_id int    `json:"ticket_id"`
	Status    string `json:"status"`
}
