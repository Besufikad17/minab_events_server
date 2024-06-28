package models

type CreateEventActionPayload struct {
	SessionVariables map[string]interface{} `json:"session_variables"`
	Input            CreateEventArgs        `json:"input"`
}

type CreateEventArgs struct {
	Title         string
	Description   string
	User_id       int
	Category_id   int
	Location_id   int
	City          string
	Venue         string
	Lat           float32
	Lng           float32
	Images        []string
	Enterance_fee float32
	Start_date    string
	End_date      string
	Tags          []string
}

type CreateEventGraphQLData struct {
	Insert_events_one CreateEventOutput `json:"insert_events_one"`
}

type CreateEventGraphQLResponse struct {
	Data   CreateEventGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError         `json:"errors,omitempty"`
}

type CreateEventOutput struct {
	Description   string  `json:"description"`
	End_date      string  `json:"end_date"`
	Enterance_fee float32 `json:"enterance_fee"`
	Id            int     `json:"id"`
	Thumbnail     string  `json:"thumbnail"`
	Start_date    string  `json:"start_date"`
	Title         string  `json:"title"`
}

type ReserveEventActionPayload struct {
	UserID int `json:"user_id"`
}
