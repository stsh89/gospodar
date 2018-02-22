package app

import (
	"app/controllers"
	"net/http"
	"os"
)

func Init() {
	controllers.RootController()
	controllers.Registrations()
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
