package controllers

import (
	"encoding/json"
	"fmt"
	"gospodar/app/config"
	"net/http"
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
		params :=
			&registrationParams{
				Email:                r.PostFormValue("email"),
				Password:             r.PostFormValue("password"),
				PasswordConfirmation: r.PostFormValue("passwordConfirmation"),
			}

		user := &user{}

		err := config.DB.QueryRow("INSERT INTO users(email, password) VALUES($1,$2) RETURNING id, email",
			params.Email, params.Password).Scan(&user.Id, &user.Email)

		if err != nil {
			panic(err)
		}

		bytes, _ := json.Marshal(user)

		fmt.Fprintf(w, string(bytes))
	})
}
