package controllers

import (
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

		json, _ := json.Marshal(params)
		fmt.Fprintf(w, string(json))
	})
}
