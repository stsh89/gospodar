package app

import (
	"gospodar/app/config"
	"gospodar/app/controllers"
	"net/http"
)

func Init() {
	config.InitGlobalVars()
	config.Database()
	controllers.RootController()
	controllers.Registrations()
	http.ListenAndServe(":"+config.Port, nil)
}
