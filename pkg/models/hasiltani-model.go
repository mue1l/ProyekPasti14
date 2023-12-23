package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mue1l/ProyekPasti14/pkg/config"
)

var db4 *gorm.DB

type Hasiltani struct {
	gorm.Model

	Nama      string `json:"nama"`
	Avatar    string `json:"avatar"`
	Deskripsi string `json:"deskripsi"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Hasiltani{})
}
func (b *Hasiltani) CreateHasiltani() *Hasiltani {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllHasiltanis() []Hasiltani {
	var Hasiltanis []Hasiltani
	db.Find(&Hasiltanis)
	return Hasiltanis
}
func GetHasiltanibyId(id int64) (*Hasiltani, *gorm.DB) {
	var getHasiltani Hasiltani
	db := db.Where("id=?", id).Find(&getHasiltani)
	return &getHasiltani, db
}
func DeleteHasiltani(id int64) Hasiltani {
	var hasiltani Hasiltani
	db.Where("id=?", id).Delete(hasiltani)
	return hasiltani
}
func UpdateHasiltani(id int64) Hasiltani {
	var hasiltani Hasiltani
	db.Where("id=?", id).Update(hasiltani)
	return hasiltani
}
