package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mue1l/ProyekPasti14/pkg/config"
)

var db8 *gorm.DB

type Galeri struct {
	gorm.Model

	Avatar    string  `json:"avatar"`
	Judul     string  `json:"judul"`
	Deskripsi string  `json:"deskripsi"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Galeri{})
}
func (b *Galeri) CreateGaleri() *Galeri {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllGaleris() []Galeri {
	var Galeris []Galeri
	db.Find(&Galeris)
	return Galeris
}
func GetGaleribyId(id int64) (*Galeri, *gorm.DB) {
	var getGaleri Galeri
	db := db.Where("id=?", id).Find(&getGaleri)
	return &getGaleri, db
}
func DeleteGaleri(id int64) Galeri {
	var galeri Galeri
	db.Where("id=?", id).Delete(galeri)
	return galeri
}
func UpdateGaleri(id int64) Galeri {
	var galeri Galeri
	db.Where("id=?", id).Update(galeri)
	return galeri
}
