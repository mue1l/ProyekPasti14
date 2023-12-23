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

var NewPerangkat models.Perangkat

func GetPerangkat(w http.ResponseWriter, r *http.Request) {
	newPerangkats := models.GetAllPerangkats()
	res, _ := json.Marshal(newPerangkats)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetPerangkatById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	perangkatId := vars["perangkatId"]
	ID, err := strconv.ParseInt(perangkatId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	perangkatDetails, _ := models.GetPerangkatbyId(ID)
	res, _ := json.Marshal(perangkatDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreatePerangkat(w http.ResponseWriter, r *http.Request) {
	CreatePerangkat := &models.Perangkat{}
	utils.ParseBody(r, CreatePerangkat)
	b := CreatePerangkat.CreatePerangkat()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeletePerangkat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	perangkatId := vars["perangkatId"]
	ID, err := strconv.ParseInt(perangkatId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	perangkat := models.DeletePerangkat(ID)
	res, _ := json.Marshal(perangkat)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdatePerangkat(w http.ResponseWriter, r *http.Request) {
	var updatePerangkat = &models.Perangkat{}
	utils.ParseBody(r, updatePerangkat)
	vars := mux.Vars(r)
	perangkatId := vars["perangkatId"]
	ID, err := strconv.ParseInt(perangkatId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	perangkatDetails, db := models.GetPerangkatbyId(ID)
	if updatePerangkat.Name != "" {
		perangkatDetails.Name = updatePerangkat.Name
	}
	if updatePerangkat.Position != "" {
		perangkatDetails.Position = updatePerangkat.Position
	}
	if updatePerangkat.Address != "" {
		perangkatDetails.Address = updatePerangkat.Address
	}
	if updatePerangkat.Avatar != "" {
		perangkatDetails.Avatar = updatePerangkat.Avatar
	}
	db.Save(&perangkatDetails)
	res, _ := json.Marshal(perangkatDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
