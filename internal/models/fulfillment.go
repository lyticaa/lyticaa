package models

import (
	"github.com/jinzhu/gorm"
)

type Fulfillment struct {
	gorm.Model
	Id   int64
	Name string
}

func GetFulfillments(db *gorm.DB) []Fulfillment {
	var fulfillments []Fulfillment
	db.Find(&fulfillments)

	return fulfillments
}
