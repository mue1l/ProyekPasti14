package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/mue1l/ProyekPasti14/pkg/controllers"
)

var AddBarang = func(router *mux.Router) {
	router.HandleFunc("/barang/", controllers.CreateBarang).Methods("POST")
	router.HandleFunc("/barang/", controllers.GetBarang).Methods("GET")
	router.HandleFunc("/barang/{barangId}", controllers.GetBarangById).Methods("GET")
	router.HandleFunc("/barang/{barangId}", controllers.UpdateBarang).Methods("PUT")
	router.HandleFunc("/barang/{barangId}", controllers.DeleteBarang).Methods("DELETE")

}