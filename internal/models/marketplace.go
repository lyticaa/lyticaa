package models

import (
	"github.com/jinzhu/gorm"
)

type Marketplace struct {
	gorm.Model
	Id   int64
	Name string
}

func GetMarketplaces(db *gorm.DB) []Marketplace {
	var marketplaces []Marketplace
	db.Find(&marketplaces)

	return marketplaces
}
