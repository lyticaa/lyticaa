package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id             int64          `db:"id"`
	UserId         string         `db:"user_id"`
	Email          string         `db:"email"`
	FirstName      sql.NullString `db:"first_name"`
	CompanyName    sql.NullString `db:"company_name"`
	SetupCompleted bool           `db:"setup_completed"`
	CreatedAt      time.Time      `db:"created_at"`
	UpdatedAt      time.Time      `db:"updated_at"`
}

func CreateUser(userId, email string, db *sqlx.DB) (*User, error) {
	user := FindUser(userId, db)
	if user.Id > 0 {
		return user, nil
	}

	r, err := db.NamedExec(`INSERT INTO users (user_id, email, setup_completed, created_at, updated_at) VALUES (:user_id, :email, :setup_completed, :created_at, :updated_at)`,
		map[string]interface{}{
			"user_id":         userId,
			"email":           email,
			"setup_completed": false,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
		})
	if err != nil {
		return user, err
	}

	user.Id, _ = r.LastInsertId()
	user.UserId = userId
	user.Email = email

	return user, nil
}

func FindUser(userId string, db *sqlx.DB) *User {
	var users []User

	err := db.Select(&users, "SELECT * FROM users WHERE user_id = $1", userId)
	if err != nil {
		logger().Error().Err(err).Msgf("unable to load the user: %v", userId)
		return &User{}
	}

	if len(users) > 0 {
		return &users[0]
	}

	return &User{}
}

func (u *User) Save(db *sqlx.DB) error {
	_, err := db.NamedExec(`UPDATE users SET email=:email, first_name=:first_name, company_name=:company_name, setup_completed=:set_completed_at WHERE user_id=:user_id`,
		map[string]interface{}{
			"user_id":         u.UserId,
			"email":           u.Email,
			"first_name":      u.FirstName,
			"company_name":    u.CompanyName,
			"setup_completed": u.SetupCompleted,
			"updated_at":      time.Now(),
		})
	if err != nil {
		return err
	}

	return nil
}
