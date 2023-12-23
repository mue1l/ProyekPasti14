package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/mue1l/ProyekPasti14/pkg/controllers"
)

var AddAgenda = func(router *mux.Router) {
	router.HandleFunc("/agenda/", controllers.CreateAgenda).Methods("POST")
	router.HandleFunc("/agenda/", controllers.GetAgenda).Methods("GET")
	router.HandleFunc("/agenda/{agendaId}", controllers.GetAgendaById).Methods("GET")
	router.HandleFunc("/agenda/{agendaId}", controllers.UpdateAgenda).Methods("PUT")
	router.HandleFunc("/agenda/{agendaId}", controllers.DeleteAgenda).Methods("DELETE")

}