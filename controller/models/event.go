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

type GetEventByIdActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            GetEventByIdArgs       `json:"input"`
}

type GetEventByIdArgs struct {
	Id int `json:"id"`
}

type GetEventByIdGraphQLRequest struct {
	Query     string           `json:"query"`
	Variables GetEventByIdArgs `json:"variables"`
}

type GetEventByIdGraphQLData struct {
	Events []GetEventByIdOutput `json:"events"`
}

type GetEventByIdGraphQLResponse struct {
	Data   GetEventByIdGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError          `json:"errors,omitempty"`
}

type GetEventByIdOutput struct {
	Description string             `json:"description"`
	End_date    string             `json:"end_date"`
	Start_date  string             `json:"start_date"`
	Title       string             `json:"title"`
	Location    CreateLocationArgs `json:"location"`
}
