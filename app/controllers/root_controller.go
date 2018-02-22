package controllers

import (
	"fmt"
	"net/http"
)

func RootController() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{ \"message\": \"Test message\" }")
	})
}
