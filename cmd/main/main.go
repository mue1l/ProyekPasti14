package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mue1l/ProyekPasti14/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.AddAgenda(r)
	routes.AddBarang(r)
	routes.AddBerita(r)
	routes.AddHasiltani(r)
	routes.AddLembaga(r)
	routes.AddPerangkat(r)
	routes.AddPengumuman(r)
	routes.AddGaleri(r)
	// routes.AddUsers(r)
	// routes.AddPendaftaran(r)
	// routes.AddPeminjaman(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
