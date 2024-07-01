package models

type AddTicketActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            AddTicketArgs          `json:"input"`
}

type AddTicketArgs struct {
	Event_id    int     `json:"event_id"`
	Ticket_type string  `json:"ticket_type"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

type AddTicketGraphQLRequest struct {
	Query     string        `json:"query"`
	Variables AddTicketArgs `json:"variables"`
}

type AddTicketGraphQLData struct {
	Insert_tickets_one AddTicketOutput `json:"insert_tickets_one"`
}

type AddTicketGraphQLResponse struct {
	Data   AddTicketGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError       `json:"errors,omitempty"`
}

type AddTicketOutput struct {
	Id int
}

type TicketInput struct {
	Description *string `json:"description"`
	Price       float32 `json:"price"`
	Ticket_type string  `json:"ticket_type"`
}
