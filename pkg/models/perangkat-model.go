package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mue1l/ProyekPasti14/pkg/config"
)

var db6 *gorm.DB

type Perangkat struct {
	gorm.Model

	Name      string `json:"name"`
	Position string `json:"position"`
	Address string `json:"address"`
	Avatar    string `json:"avatar"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Perangkat{})
}
func (b *Perangkat) CreatePerangkat() *Perangkat {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllPerangkats() []Perangkat {
	var Perangkats []Perangkat
	db.Find(&Perangkats)
	return Perangkats
}
func GetPerangkatbyId(id int64) (*Perangkat, *gorm.DB) {
	var getPerangkat Perangkat
	db := db.Where("id=?", id).Find(&getPerangkat)
	return &getPerangkat, db
}
func DeletePerangkat(id int64) Perangkat {
	var perangkat Perangkat
	db.Where("id=?", id).Delete(perangkat)
	return perangkat
}
func UpdatePerangkat(id int64) Perangkat {
	var perangkat Perangkat
	db.Where("id=?", id).Update(perangkat)
	return perangkat
}
