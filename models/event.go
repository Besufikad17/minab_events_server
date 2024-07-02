package models

type CreateEventActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            CreateEventArgs        `json:"input"`
}

type CreateEventArgs struct {
	Title       string
	Description string
	User_id     int
	Category_id int
	Location_id int
	City        string
	Venue       string
	Lat         float32
	Lng         float32
	Images      []string
	Tickets     []TicketInput
	Start_date  string
	End_date    string
	Tags        []string
}

type CreateEventGraphQLData struct {
	Insert_events_one CreateEventOutput `json:"insert_events_one"`
}

type CreateEventGraphQLResponse struct {
	Data   CreateEventGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError         `json:"errors,omitempty"`
}

type CreateEventOutput struct {
	Description string `json:"description"`
	End_date    string `json:"end_date"`
	Id          int    `json:"id"`
	Thumbnail   string `json:"thumbnail"`
	Start_date  string `json:"start_date"`
	Title       string `json:"title"`
}

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
