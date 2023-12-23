package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/mue1l/ProyekPasti14/pkg/controllers"
)

var AddGaleri = func(router *mux.Router) {
	router.HandleFunc("/galeri/", controllers.CreateGaleri).Methods("POST")
	router.HandleFunc("/galeri/", controllers.GetGaleri).Methods("GET")
	router.HandleFunc("/galeri/{galeriId}", controllers.GetGaleriById).Methods("GET")
	router.HandleFunc("/galeri/{galeriId}", controllers.UpdateGaleri).Methods("PUT")
	router.HandleFunc("/galeri/{galeriId}", controllers.DeleteGaleri).Methods("DELETE")

}