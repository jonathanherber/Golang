package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanherber/Golang_api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter() //instacia do gorilla mux
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
