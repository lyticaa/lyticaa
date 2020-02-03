package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id               int64
	UserId           string `db:"user_id"`
	Email            string
	FirstName        string    `db:"first_name"`
	CompanyName      string    `db:"company_name"`
	SetupCompletedAt time.Time `db:"setup_completed_at"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

func CreateUser(userId, email string, db *sqlx.DB) (User, error) {
	user := FindUser(userId, db)
	if user.Id == int64(0) {
		query := `INSERT INTO users (user_id, email, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`

		r, err := db.Exec(query, userId, email)
		if err != nil {
			logger().Error().Err(err)
			return user, err
		}

		user.Id, _ = r.LastInsertId()
		user.UserId = userId
		user.Email = email
	}

	return user, nil
}

func UpdateUser(userId, email, firstName, companyName string, db *sqlx.DB) error {
	query := `UPDATE users SET email=?, first_name=?, company_name=? WHERE user_id=?`

	_, err := db.Exec(query, email, firstName, companyName, userId)
	if err != nil {
		logger().Error().Err(err)
		return err
	}

	return nil
}

func FindUser(userId string, db *sqlx.DB) User {
	user := User{}

	err := db.Get(&user, "SELECT id, user_id, email, fist_name, company_name, setup_completed_at, created_at FROM users WHERE user_id=$1", userId)
	if err != nil {
		logger().Error().Err(err)
	}

	return user
}

func SetupCompleted(userId string, setupCompletedAt time.Time, db *sqlx.DB) error {
	query := `UPDATE users SET setup_completed_at=? WHERE user_id=?`

	_, err := db.Exec(query, setupCompletedAt, userId)
	if err != nil {
		logger().Error().Err(err)
		return err
	}

	return nil
}
