package app

import (
	"mvc/controllers"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
