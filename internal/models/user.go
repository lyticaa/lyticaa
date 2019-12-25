package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserId string
	Email  string
}

func CreateUser(userId, email string, db *gorm.DB) {
	user := FindUserByUserId(userId, db)
	if user.UserId == "" {
		db.Create(&User{UserId: userId, Email: email})
	} else {
		UpdateUser(user, db)
	}
}

func FindUserByUserId(userId string, db *gorm.DB) User {
	var user User
	db.First(&user, "user_id = ?", userId)

	return user
}

func UpdateUser(user User, db *gorm.DB) {
	db.Model(&user).Update("updated_at", time.Now())
}
