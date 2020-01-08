package models

import (
	"github.com/jinzhu/gorm"
)

type TaxCollectionModel struct {
	gorm.Model
	Id   int64
	Name string
}

func GetTaxCollectionModels(db *gorm.DB) []TaxCollectionModel {
	var collections []TaxCollectionModel
	db.Find(&collections)

	return collections
}
