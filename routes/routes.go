package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jonathanherber/Golang_api/controllers"
	"github.com/jonathanherber/Golang_api/middleware"
)

func HandleRequest() {
	r := mux.NewRouter() //instancia do gorilla mux, que é um router de url
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaUmaPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
