package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mue1l/ProyekPasti14/pkg/config"
)

var db3 *gorm.DB

type Berita struct {
	gorm.Model

	Avatar    string  `json:"avatar"`
	Judul     string  `json:"judul"`
	Deskripsi string  `json:"deskripsi"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Berita{})
}
func (b *Berita) CreateBerita() *Berita {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBeritas() []Berita {
	var Beritas []Berita
	db.Find(&Beritas)
	return Beritas
}
func GetBeritabyId(id int64) (*Berita, *gorm.DB) {
	var getBerita Berita
	db := db.Where("id=?", id).Find(&getBerita)
	return &getBerita, db
}
func DeleteBerita(id int64) Berita {
	var berita Berita
	db.Where("id=?", id).Delete(berita)
	return berita
}
func UpdateBerita(id int64) Berita {
	var berita Berita
	db.Where("id=?", id).Update(berita)
	return berita
}
