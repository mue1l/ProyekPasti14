package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mue1l/ProyekPasti14/pkg/config"
)

var db7 *gorm.DB

type Pengumuman struct {
	gorm.Model
	Avatar    string `json:"avatar"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Pengumuman{})
}
func (b *Pengumuman) CreatePengumuman() *Pengumuman {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllPengumumans() []Pengumuman {
	var Pengumumans []Pengumuman
	db.Find(&Pengumumans)
	return Pengumumans
}
func GetPengumumanbyId(id int64) (*Pengumuman, *gorm.DB) {
	var getPengumuman Pengumuman
	db := db.Where("id=?", id).Find(&getPengumuman)
	return &getPengumuman, db
}
func DeletePengumuman(id int64) Pengumuman {
	var Pengumuman Pengumuman
	db.Where("id=?", id).Delete(Pengumuman)
	return Pengumuman
}
func UpdatePengumuman(id int64) Pengumuman {
	var Pengumuman Pengumuman
	db.Where("id=?", id).Update(Pengumuman)
	return Pengumuman
}
