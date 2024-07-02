package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"os"
	"text/template"

	"github.com/Besufikad17/minab_events/hasura/actions"
	"github.com/Besufikad17/minab_events/models"
)

func NotifyUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var actionPayload models.ReserveEventActionPayload
	err = json.Unmarshal(reqBody, &actionPayload)
	if err != nil {
		println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		Name string
	}{
		Name: user.First_name + " " + user.Last_name,
	})

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
