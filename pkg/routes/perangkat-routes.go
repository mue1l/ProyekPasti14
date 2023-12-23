package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/mue1l/ProyekPasti14/pkg/controllers"
)

var AddPerangkat= func(router *mux.Router) {
	router.HandleFunc("/perangkat/", controllers.CreatePerangkat).Methods("POST")
	router.HandleFunc("/perangkat/", controllers.GetPerangkat).Methods("GET")
	router.HandleFunc("/perangkat/{perangkatId}", controllers.GetPerangkatById).Methods("GET")
	router.HandleFunc("/perangkat/{perangkatId}", controllers.UpdatePerangkat).Methods("PUT")
	router.HandleFunc("/perangkat/{perangkatId}", controllers.DeletePerangkat).Methods("DELETE")

}