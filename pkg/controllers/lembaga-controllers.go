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

var NewLembaga models.Lembaga

func GetLembaga(w http.ResponseWriter, r *http.Request) {
	newLembagas := models.GetAllLembagas()
	res, _ := json.Marshal(newLembagas)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetLembagaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lembagaId := vars["lembagaId"]
	ID, err := strconv.ParseInt(lembagaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	lembagaDetails, _ := models.GetLembagabyId(ID)
	res, _ := json.Marshal(lembagaDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateLembaga(w http.ResponseWriter, r *http.Request) {
	CreateLembaga := &models.Lembaga{}
	utils.ParseBody(r, CreateLembaga)
	b := CreateLembaga.CreateLembaga()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteLembaga(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lembagaId := vars["lembagaId"]
	ID, err := strconv.ParseInt(lembagaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	lembaga := models.DeleteLembaga(ID)
	res, _ := json.Marshal(lembaga)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateLembaga(w http.ResponseWriter, r *http.Request) {
	var updateLembaga = &models.Lembaga{}
	utils.ParseBody(r, updateLembaga)
	vars := mux.Vars(r)
	lembagaId := vars["lembagaId"]
	ID, err := strconv.ParseInt(lembagaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	lembagaDetails, db := models.GetLembagabyId(ID)
	if updateLembaga.Name != "" {
		lembagaDetails.Name = updateLembaga.Name
	}
	if updateLembaga.Position != "" {
		lembagaDetails.Position = updateLembaga.Position
	}
	if updateLembaga.Address != "" {
		lembagaDetails.Address = updateLembaga.Address
	}
	if updateLembaga.Avatar != "" {
		lembagaDetails.Avatar = updateLembaga.Avatar
	}
	db.Save(&lembagaDetails)
	res, _ := json.Marshal(lembagaDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
