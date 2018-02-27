package controllers

import (
	"encoding/json"
	"fmt"
	"gospodar/app/config"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"net/smtp"
	"os"
)

type registrationParams struct {
	Email                string
	Password             string
	PasswordConfirmation string
}

type user struct {
	Id    int
	Email string
}

func Registrations() {
	http.HandleFunc("/registrations", func(w http.ResponseWriter, r *http.Request) {
		params := &registrationParams{
			Email:                r.PostFormValue("email"),
			Password:             r.PostFormValue("password"),
			PasswordConfirmation: r.PostFormValue("passwordConfirmation"),
		}

		user := &user{}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)

		err := config.DB.QueryRow("INSERT INTO users(email, password) VALUES($1,$2) RETURNING id, email",
			params.Email, string(hashedPassword)).Scan(&user.Id, &user.Email)

		if err != nil {
			panic(err)
		}

		bytes, _ := json.Marshal(user)
		sendRegistrationEmail(user.Email)
		fmt.Fprintf(w, string(bytes))
	})
}

func sendRegistrationEmail(emailAddress string) {
	from := os.Getenv("SMTP_USER_NAME")
	password := os.Getenv("SMTP_PASSWORD")
	domain := os.Getenv("SMTP_DOMAIN")
	port := os.Getenv("SMTP_PORT")
	msg := "From: " + from + "\n" +
		"To: " + emailAddress + "\n" +
		"Subject: Registration\n\n" +
		"Registration Complete!"

	fmt.Println(from, password, domain, port, msg)
	err := smtp.SendMail(
		domain + ":" + port,
		smtp.PlainAuth("", from, password, domain),
		from,
		[]string{emailAddress},
		[]byte(msg),
	)

	if err != nil {
		panic(err)
	}
}
