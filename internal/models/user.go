package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id        int64
	UserId    string `db:"user_id"`
	Email     string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func CreateUser(userId, email string, db *sqlx.DB) {
	user := FindUserByUserId(userId, db)
	if user.UserId == "" {
		query := `INSERT INTO users (user_id, email, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`
		db.Exec(query, userId, email)
	}
}

func FindUserByUserId(userId string, db *sqlx.DB) User {
	user := User{}
	err := db.Get(&user, "SELECT id, user_id, email, created_at FROM users WHERE user_id=$1", userId)
	if err != nil {
		panic(err)
	}

	return user
}
