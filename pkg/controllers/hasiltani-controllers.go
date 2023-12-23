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

var NewHasiltani models.Hasiltani

func GetHasiltani(w http.ResponseWriter, r *http.Request) {
	newHasiltanis := models.GetAllHasiltanis()
	res, _ := json.Marshal(newHasiltanis)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetHasiltaniById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hasiltaniId := vars["hasiltaniId"]
	ID, err := strconv.ParseInt(hasiltaniId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	hasiltaniDetails, _ := models.GetHasiltanibyId(ID)
	res, _ := json.Marshal(hasiltaniDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateHasiltani(w http.ResponseWriter, r *http.Request) {
	CreateHasiltani := &models.Hasiltani{}
	utils.ParseBody(r, CreateHasiltani)
	b := CreateHasiltani.CreateHasiltani()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteHasiltani(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hasiltaniId := vars["hasiltaniId"]
	ID, err := strconv.ParseInt(hasiltaniId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	hasiltani := models.DeleteHasiltani(ID)
	res, _ := json.Marshal(hasiltani)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateHasiltani(w http.ResponseWriter, r *http.Request) {
	var updateHasiltani = &models.Hasiltani{}
	utils.ParseBody(r, updateHasiltani)
	vars := mux.Vars(r)
	hasiltaniId := vars["hasiltaniId"]
	ID, err := strconv.ParseInt(hasiltaniId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	hasiltaniDetails, db := models.GetHasiltanibyId(ID)
	if updateHasiltani.Nama != "" {
		hasiltaniDetails.Nama = updateHasiltani.Nama
	}
	if updateHasiltani.Avatar != "" {
		hasiltaniDetails.Avatar = updateHasiltani.Avatar
	}
	if updateHasiltani.Deskripsi != "" {
		hasiltaniDetails.Deskripsi = updateHasiltani.Deskripsi
	}
	db.Save(&hasiltaniDetails)
	res, _ := json.Marshal(hasiltaniDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
