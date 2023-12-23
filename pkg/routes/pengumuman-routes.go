package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/mue1l/ProyekPasti14/pkg/controllers"
)

var AddPengumuman= func(router *mux.Router) {
	router.HandleFunc("/pengumuman/", controllers.CreatePengumuman).Methods("POST")
	router.HandleFunc("/pengumuman/", controllers.GetPengumuman).Methods("GET")
	router.HandleFunc("/pengumuman/{pengumumanId}", controllers.GetPengumumanById).Methods("GET")
	router.HandleFunc("/pengumuman/{pengumumanId}", controllers.UpdatePengumuman).Methods("PUT")
	router.HandleFunc("/pengumuman/{pengumumanId}", controllers.DeletePengumuman).Methods("DELETE")

}