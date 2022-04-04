package routes

import (
	"log"
	"net/http"

	"github.com/jonathanherber/Golang_api/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
