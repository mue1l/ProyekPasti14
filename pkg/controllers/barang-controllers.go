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

var NewBarang models.Barang

func GetBarang(w http.ResponseWriter, r *http.Request) {
	newBarangs := models.GetAllBarangs()
	res, _ := json.Marshal(newBarangs)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBarangById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barangId := vars["barangId"]
	ID, err := strconv.ParseInt(barangId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	barangDetails, _ := models.GetBarangbyId(ID)
	res, _ := json.Marshal(barangDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBarang(w http.ResponseWriter, r *http.Request) {
	CreateBarang := &models.Barang{}
	utils.ParseBody(r, CreateBarang)
	b := CreateBarang.CreateBarang()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBarang(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barangId := vars["barangId"]
	ID, err := strconv.ParseInt(barangId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	barang := models.DeleteBarang(ID)
	res, _ := json.Marshal(barang)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBarang(w http.ResponseWriter, r *http.Request) {
	var updateBarang = &models.Barang{}
	utils.ParseBody(r, updateBarang)
	vars := mux.Vars(r)
	barangId := vars["barangId"]
	ID, err := strconv.ParseInt(barangId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	barangDetails, db := models.GetBarangbyId(ID)
	if updateBarang.Nama != "" {
		barangDetails.Nama = updateBarang.Nama
	}
	if updateBarang.Jumlah != "" {
		barangDetails.Jumlah = updateBarang.Jumlah
	}
	if updateBarang.Avatar != "" {
		barangDetails.Avatar = updateBarang.Avatar
	}
	db.Save(&barangDetails)
	res, _ := json.Marshal(barangDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
