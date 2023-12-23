package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/mue1l/ProyekPasti14/pkg/controllers"
)

var AddLembaga= func(router *mux.Router) {
	router.HandleFunc("/lembaga/", controllers.CreateLembaga).Methods("POST")
	router.HandleFunc("/lembaga/", controllers.GetLembaga).Methods("GET")
	router.HandleFunc("/lembaga/{lembagaId}", controllers.GetLembagaById).Methods("GET")
	router.HandleFunc("/lembaga/{lembagaId}", controllers.UpdateLembaga).Methods("PUT")
	router.HandleFunc("/lembaga/{lembagaId}", controllers.DeleteLembaga).Methods("DELETE")

}