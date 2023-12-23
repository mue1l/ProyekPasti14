package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mue1l/ProyekPasti14/pkg/config"
)

var db *gorm.DB

type Agenda struct {
	gorm.Model
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Agenda{})
}
func (b *Agenda) CreateAgenda() *Agenda {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllAgendas() []Agenda {
	var Agendas []Agenda
	db.Find(&Agendas)
	return Agendas
}
func GetAgendabyId(id int64) (*Agenda, *gorm.DB) {
	var getAgenda Agenda
	db := db.Where("id=?", id).Find(&getAgenda)
	return &getAgenda, db
}
func DeleteAgenda(id int64) Agenda {
	var agenda Agenda
	db.Where("id=?", id).Delete(agenda)
	return agenda
}
func UpdateAgenda(id int64) Agenda {
	var agenda Agenda
	db.Where("id=?", id).Update(agenda)
	return agenda
}
