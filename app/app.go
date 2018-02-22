package app

import (
	"app/controllers"
	"net/http"
	"os"
)

func Init() {
	controllers.RootController()
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
