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

var NewGaleri models.Galeri

func GetGaleri(w http.ResponseWriter, r *http.Request) {
	newGaleris := models.GetAllGaleris()
	res, _ := json.Marshal(newGaleris)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetGaleriById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lembagaId := vars["lembagaId"]
	ID, err := strconv.ParseInt(lembagaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	lembagaDetails, _ := models.GetGaleribyId(ID)
	res, _ := json.Marshal(lembagaDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateGaleri(w http.ResponseWriter, r *http.Request) {
	CreateGaleri := &models.Galeri{}
	utils.ParseBody(r, CreateGaleri)
	b := CreateGaleri.CreateGaleri()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteGaleri(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	galeriId := vars["galeriId"]
	ID, err := strconv.ParseInt(galeriId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	galeri := models.DeleteGaleri(ID)
	res, _ := json.Marshal(galeri)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateGaleri(w http.ResponseWriter, r *http.Request) {
	var UpdateGaleri = &models.Galeri{}
	utils.ParseBody(r, UpdateGaleri)
	vars := mux.Vars(r)
	galeriId := vars["galeriId"]
	ID, err := strconv.ParseInt(galeriId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	galeriDetails, db := models.GetGaleribyId(ID)
	if UpdateGaleri.Avatar != "" {
		galeriDetails.Avatar = UpdateGaleri.Avatar
	}
	if UpdateGaleri.Judul != "" {
		galeriDetails.Judul = UpdateGaleri.Judul
	}
	if UpdateGaleri.Deskripsi != "" {
		galeriDetails.Deskripsi = UpdateGaleri.Deskripsi
	}
	db.Save(&galeriDetails)
	res, _ := json.Marshal(galeriDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
