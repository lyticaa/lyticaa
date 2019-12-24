package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId string
	Email  string
}
