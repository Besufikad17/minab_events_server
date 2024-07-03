package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"text/template"

	"github.com/Besufikad17/minab_events/hasura/actions"
	"github.com/Besufikad17/minab_events/models"
	"github.com/Besufikad17/minab_events/utils/helpers"
)

func NotifyUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var actionPayload models.OnEventReservedActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// getting user info
	user, err := actions.GetUserById(map[string]interface{}{
		"id": actionPayload.Input.User_id,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// getting ticket info
	ticket, err := actions.GetTicketById(models.GetTicketByIdArgs{
		Id: actionPayload.Input.Ticket_id,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// getting event info
	event, err := actions.GetEventById(models.GetEventByIdArgs{
		Id: actionPayload.Input.Event_id,
	})

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	qrCode, err := helpers.GenerateQR(strconv.Itoa(actionPayload.Input.Id))

	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PWD")

	to := []string{
		user.Email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	t, err := template.ParseFiles("public/templates/email.html")

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Event reserved \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Description string
		FirstName   string
		LastName    string
		QrCode      string
		StartDate   string
		Title       string
		Type        string
	}{
		Description: event.Description,
		FirstName:   user.First_name,
		LastName:    user.Last_name,
		QrCode:      *qrCode,
		StartDate:   event.Start_date,
		Title:       event.Title,
		Type:        ticket.Ticket_type,
	})

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
