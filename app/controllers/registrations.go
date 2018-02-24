package controllers

import (
	"gospodar/app/config"
	"encoding/json"
	"fmt"
	"net/http"
)

type registrationParams struct {
	Email                string
	Password             string
	PasswordConfirmation string
}

func Registrations() {
	http.HandleFunc("/registrations", func(w http.ResponseWriter, r *http.Request) {
		params :=
			&registrationParams{
				Email:                r.PostFormValue("email"),
				Password:             r.PostFormValue("password"),
				PasswordConfirmation: r.PostFormValue("passwordConfirmation"),
			}

		email := ""
		json, _ := json.Marshal(params)
		fmt.Println(string(json))
		err := config.DB.QueryRow("INSERT INTO users(email, password) VALUES($1,$2) RETURNING email", params.Email, params.Password).Scan(&email)

		if err != nil {
	    panic(err)
	  }

		fmt.Fprintf(w, email)
	})
}
