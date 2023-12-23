package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/mue1l/ProyekPasti14/pkg/controllers"
)

var AddBerita = func(router *mux.Router) {
	router.HandleFunc("/berita/", controllers.CreateBerita).Methods("POST")
	router.HandleFunc("/berita/", controllers.GetBerita).Methods("GET")
	router.HandleFunc("/berita/{beritaId}", controllers.GetBeritaById).Methods("GET")
	router.HandleFunc("/berita/{beritaId}", controllers.UpdateBerita).Methods("PUT")
	router.HandleFunc("/berita/{beritaId}", controllers.DeleteBerita).Methods("DELETE")

}