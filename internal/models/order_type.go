package models

import (
	"github.com/jinzhu/gorm"
)

type OrderType struct {
	gorm.Model
	Id   int64
	Name string
}

func GetOrderTypes(db *gorm.DB) []OrderType {
	var orderTypes []OrderType
	db.Find(&orderTypes)

	return orderTypes
}
