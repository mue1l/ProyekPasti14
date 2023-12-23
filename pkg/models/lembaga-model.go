package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mue1l/ProyekPasti14/pkg/config"
)

var db5 *gorm.DB

type Lembaga struct {
	gorm.Model

	Name      string `json:"name"`
	Position string `json:"position"`
	Address string `json:"address"`
	Avatar    string `json:"avatar"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Lembaga{})
}
func (b *Lembaga) CreateLembaga() *Lembaga {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllLembagas() []Lembaga {
	var Lembagas []Lembaga
	db.Find(&Lembagas)
	return Lembagas
}
func GetLembagabyId(id int64) (*Lembaga, *gorm.DB) {
	var getLembaga Lembaga
	db := db.Where("id=?", id).Find(&getLembaga)
	return &getLembaga, db
}
func DeleteLembaga(id int64) Lembaga {
	var lembaga Lembaga
	db.Where("id=?", id).Delete(lembaga)
	return lembaga
}
func UpdateLembaga(id int64) Lembaga {
	var lembaga Lembaga
	db.Where("id=?", id).Update(lembaga)
	return lembaga
}
