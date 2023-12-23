package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/mue1l/ProyekPasti14/pkg/controllers"
)

var AddHasiltani= func(router *mux.Router) {
	router.HandleFunc("/hasiltani/", controllers.CreateHasiltani).Methods("POST")
	router.HandleFunc("/hasiltani/", controllers.GetHasiltani).Methods("GET")
	router.HandleFunc("/hasiltani/{hasiltaniId}", controllers.GetHasiltaniById).Methods("GET")
	router.HandleFunc("/hasiltani/{hasiltaniId}", controllers.UpdateHasiltani).Methods("PUT")
	router.HandleFunc("/hasiltani/{hasiltaniId}", controllers.DeleteHasiltani).Methods("DELETE")

}