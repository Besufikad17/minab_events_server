package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	actions "github.com/Besufikad17/minab_events/hasura/actions"
	models "github.com/Besufikad17/minab_events/models"
	helpers "github.com/Besufikad17/minab_events/utils/helpers"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var actionPayload models.CreateEventActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	location, err := actions.CreateLocation(models.CreateLocationArgs{
		City:  actionPayload.Input.City,
		Venue: actionPayload.Input.Venue,
		Lat:   actionPayload.Input.Lat,
		Lng:   actionPayload.Input.Lng,
	}, r.Header.Get("Authorization"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	thumbnail, err := helpers.SaveImageToFile(actionPayload.Input.Images[0])
	mapVariables := map[string]interface{}{
		"title":       actionPayload.Input.Title,
		"description": actionPayload.Input.Description,
		"thumbnail":   "http://localhost:5000/images/" + *thumbnail,
		"user_id":     actionPayload.Input.User_id,
		"category_id": actionPayload.Input.Category_id,
		"location_id": location.Id,
		"start_date":  actionPayload.Input.Start_date,
		"end_date":    actionPayload.Input.End_date,
	}

	result, err := actions.CreateEvent(mapVariables, r.Header.Get("Authorization"))

	if err != nil {
		errorObject := models.GraphQLError{
			Message: err.Error(),
		}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}

	for _, image := range actionPayload.Input.Images {
		fileName, err := helpers.SaveImageToFile(image)

		if err != nil {
			errorObject := models.GraphQLError{
				Message: err.Error(),
			}
			errorBody, _ := json.Marshal(errorObject)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorBody)
			return
		}

		_, err = actions.CreateImage(models.CreateImageArgs{
			Event_id: result.Id,
			Url:      "http://localhost:5000/images/" + *fileName,
		}, r.Header.Get("Authorization"))

		if err != nil {
			errorObject := models.GraphQLError{
				Message: err.Error(),
			}
			errorBody, _ := json.Marshal(errorObject)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorBody)
			return
		}
	}

	for _, tag := range actionPayload.Input.Tags {
		_, err = actions.CreateTag(models.CreateTagArgs{
			Name:     tag,
			Event_id: result.Id,
		}, r.Header.Get("Authorization"))

		if err != nil {
			println(err)
			errorObject := models.GraphQLError{
				Message: err.Error(),
			}
			errorBody, _ := json.Marshal(errorObject)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorBody)
			return
		}
	}

	for _, ticket := range actionPayload.Input.Tickets {
		_, err = actions.AddTicket(models.AddTicketArgs{
			Event_id:    result.Id,
			Ticket_type: ticket.Ticket_type,
			Description: *ticket.Description,
			Price:       ticket.Price,
		}, r.Header.Get("Authorization"))

		if err != nil {
			println(err)
			errorObject := models.GraphQLError{
				Message: err.Error(),
			}
			errorBody, _ := json.Marshal(errorObject)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(errorBody)
			return
		}
	}

	data, _ := json.Marshal(result)
	w.Write(data)
}

func ReserveEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	var actionPayload models.ReserveEventActionPayload
	actionPayload.Input.Status = "pending"
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	user, err := actions.GetUserById(map[string]interface{}{
		"id": actionPayload.Input.User_id,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ticket, err := actions.GetTicketById(models.GetTicketByIdArgs{
		Id: actionPayload.Input.Ticket_id,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	textRef := "MinabEvents" + strconv.FormatInt(time.Now().UnixMilli(), 10) + user.Phone_number
	payment, err := actions.CreatePayment(models.CreatePaymentArgs{
		User_id:   actionPayload.Input.User_id,
		Ticket_id: actionPayload.Input.Ticket_id,
		Amount:    ticket.Price,
		Reference: textRef,
		Status:    "pending",
	}, r.Header.Get("Authorization"))

	callbackUrl := os.Getenv("CHAPA_CALLBACK_URL")
	returnUrl := os.Getenv("CHAPA_RETURN_URL")
	chapaArgs := models.ChapaArgs{
		Amount:       ticket.Price,
		Currency:     "ETB",
		First_name:   user.First_name,
		Last_name:    user.Last_name,
		Email:        user.Email,
		Phone_number: user.Phone_number,
		Tx_ref:       textRef,
		Return_url:   returnUrl + "?pid=" + strconv.Itoa(payment.Id),
		Callback_url: callbackUrl,
	}

	chapaResult, err := actions.PayWithChapa(chapaArgs)

	result, err := actions.ReserveEvent(actionPayload.Input, r.Header.Get("Authorization"))
	if result != nil {
		result.CheckoutUrl = chapaResult.Data.Checkout_url
	}

	if err != nil {
		errorObject := models.GraphQLError{
			Message: err.Error(),
		}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}

	data, _ := json.Marshal(result)
	w.Write(data)
}
