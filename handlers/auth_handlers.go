package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	actions "github.com/Besufikad17/minab_events/hasura/actions"
	models "github.com/Besufikad17/minab_events/models"
	helpers "github.com/Besufikad17/minab_events/utils/helpers"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	var actionPayload models.RegisterActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	result, err := actions.Register(actionPayload.Input)
	if err != nil {
		errorObject := models.GraphQLError{
			Message: err.Error(),
		}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}

	if result == nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	token, err := helpers.CreateToken(models.User{
		ID:          result.Id,
		FirstName:   *result.First_name,
		LastName:    *result.Last_name,
		Email:       *result.Email,
		PhoneNumber: *result.Phone_number,
	}, actionPayload.Input.Remember_me)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, _ := json.Marshal(
		map[string]interface{}{
			"token": token,
		},
	)
	w.Write(data)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	var actionPayload models.LoginActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	result, err := actions.SearchUser(models.SearchUserArgs{
		Login_text: actionPayload.Input.Login_text,
	})

	if err != nil {
		errorObject := models.GraphQLError{
			Message: err.Error(),
		}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}

	if helpers.Compare(result.Password, actionPayload.Input.Password) {
		token, err := helpers.CreateToken(models.User{
			ID:          result.Id,
			FirstName:   result.First_name,
			LastName:    result.Last_name,
			Email:       result.Email,
			PhoneNumber: result.Phone_number,
		}, actionPayload.Input.Remember_me)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data, _ := json.Marshal(
			map[string]interface{}{
				"token": token,
			},
		)
		w.Write(data)
	} else {
		errorObject := models.GraphQLError{
			Message: "Invalid credentials!!",
		}
		errorBody, _ := json.Marshal(errorObject)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorBody)
		return
	}
}
