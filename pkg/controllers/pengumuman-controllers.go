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

var NewPengumuman models.Pengumuman

func GetPengumuman(w http.ResponseWriter, r *http.Request) {
	newPengumumans := models.GetAllPengumumans()
	res, _ := json.Marshal(newPengumumans)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetPengumumanById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lembagaId := vars["lembagaId"]
	ID, err := strconv.ParseInt(lembagaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	lembagaDetails, _ := models.GetPengumumanbyId(ID)
	res, _ := json.Marshal(lembagaDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreatePengumuman(w http.ResponseWriter, r *http.Request) {
	CreatePengumuman := &models.Pengumuman{}
	utils.ParseBody(r, CreatePengumuman)
	b := CreatePengumuman.CreatePengumuman()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeletePengumuman(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pengumumanId := vars["pengumumanId"]
	ID, err := strconv.ParseInt(pengumumanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pengumuman := models.DeletePengumuman(ID)
	res, _ := json.Marshal(pengumuman)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePengumuman(w http.ResponseWriter, r *http.Request) {
	var updatePengumuman = &models.Pengumuman{}
	utils.ParseBody(r, updatePengumuman)
	vars := mux.Vars(r)
	pengumumanId := vars["pengumumanId"]
	ID, err := strconv.ParseInt(pengumumanId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	pengumumanDetails, db := models.GetPengumumanbyId(ID)
	if updatePengumuman.Avatar != "" {
		pengumumanDetails.Avatar = updatePengumuman.Avatar
	}
	if updatePengumuman.Judul != "" {
		pengumumanDetails.Judul = updatePengumuman.Judul
	}
	if updatePengumuman.Deskripsi != "" {
		pengumumanDetails.Deskripsi = updatePengumuman.Deskripsi
	}
	db.Save(&pengumumanDetails)
	res, _ := json.Marshal(pengumumanDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
