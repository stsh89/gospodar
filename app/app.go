package app

import (
	"gospodar/app/config"
	"gospodar/app/controllers"
	"net/http"
	"os"
)

func Init() {
	config.Database()
	controllers.RootController()
	controllers.Registrations()
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
