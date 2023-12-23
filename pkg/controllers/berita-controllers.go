package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mue1l/ProyekPasti14/pkg/models"
	"github.com/mue1l/ProyekPasti14/pkg/utils"
)

var NewBerita models.Berita

func GetBerita(w http.ResponseWriter, r *http.Request) {
	newBeritas := models.GetAllBeritas()
	res, _ := json.Marshal(newBeritas)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBeritaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	beritaId := vars["beritaId"]
	ID, err := strconv.ParseInt(beritaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	beritaDetails, _ := models.GetBeritabyId(ID)
	res, _ := json.Marshal(beritaDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBerita(w http.ResponseWriter, r *http.Request) {
	CreateBerita := &models.Berita{}
	utils.ParseBody(r, CreateBerita)
	b := CreateBerita.CreateBerita()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBerita(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	beritaId := vars["beritaId"]
	ID, err := strconv.ParseInt(beritaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	berita := models.DeleteBerita(ID)
	res, _ := json.Marshal(berita)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBerita(w http.ResponseWriter, r *http.Request) {
	var updateBerita = &models.Berita{}
	utils.ParseBody(r, updateBerita)
	vars := mux.Vars(r)
	beritaId := vars["beritaId"]
	ID, err := strconv.ParseInt(beritaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	beritaDetails, db := models.GetBeritabyId(ID)
	if updateBerita.Avatar != "" {
		beritaDetails.Avatar = updateBerita.Avatar
	}
	if updateBerita.Judul != "" {
		beritaDetails.Judul = updateBerita.Judul
	}
	if updateBerita.Deskripsi != "" {
		beritaDetails.Deskripsi = updateBerita.Deskripsi
	}
	db.Save(&beritaDetails)
	res, _ := json.Marshal(beritaDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
