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

var NewAgenda models.Agenda

func GetAgenda(w http.ResponseWriter, r *http.Request) {
	newAgendas := models.GetAllAgendas()
	res, _ := json.Marshal(newAgendas)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetAgendaById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	agendaId := vars["agendaId"]
	ID, err := strconv.ParseInt(agendaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	agendaDetails, _ := models.GetAgendabyId(ID)
	res, _ := json.Marshal(agendaDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateAgenda(w http.ResponseWriter, r *http.Request) {
	CreateAgenda := &models.Agenda{}
	utils.ParseBody(r, CreateAgenda)
	b := CreateAgenda.CreateAgenda()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteAgenda(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	agendaId := vars["agendaId"]
	ID, err := strconv.ParseInt(agendaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	agenda := models.DeleteAgenda(ID)
	res, _ := json.Marshal(agenda)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAgenda(w http.ResponseWriter, r *http.Request) {
	var updateAgenda = &models.Agenda{}
	utils.ParseBody(r, updateAgenda)
	vars := mux.Vars(r)
	agendaId := vars["agendaId"]
	ID, err := strconv.ParseInt(agendaId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	agendaDetails, db := models.GetAgendabyId(ID)
	if updateAgenda.Judul != "" {
		agendaDetails.Judul = updateAgenda.Judul
	}
	if updateAgenda.Deskripsi != "" {
		agendaDetails.Deskripsi = updateAgenda.Deskripsi
	}
	db.Save(&agendaDetails)
	res, _ := json.Marshal(agendaDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
