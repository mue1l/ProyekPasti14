package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mue1l/ProyekPasti14/pkg/config"
)

var db2 *gorm.DB

type Barang struct {
	gorm.Model
	Nama   string `json:"nama"`
	Jumlah string `json:"jumlah"`
	Avatar string `json:"avatar"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Barang{})
}
func (b *Barang) CreateBarang() *Barang {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBarangs() []Barang {
	var Barangs []Barang
	db.Find(&Barangs)
	return Barangs
}
func GetBarangbyId(id int64) (*Barang, *gorm.DB) {
	var getBarang Barang
	db := db.Where("id=?", id).Find(&getBarang)
	return &getBarang, db
}
func DeleteBarang(id int64) Barang {
	var barang Barang
	db.Where("id=?", id).Delete(barang)
	return barang
}
func UpdateBarang(id int64) Barang {
	var barang Barang
	db.Where("id=?", id).Update(barang)
	return barang
}
