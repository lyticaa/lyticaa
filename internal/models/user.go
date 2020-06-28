package models

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type User struct {
	Id                   int64          `db:"id"`
	UserId               string         `db:"user_id"`
	Email                string         `db:"email"`
	StripeUserId         sql.NullString `db:"stripe_user_id"`
	StripeSubscriptionId sql.NullString `db:"stripe_subscription_id"`
	StripePlanId         sql.NullString `db:"stripe_plan_id"`
	Nickname             string
	Picture              string
	SetupCompleted       bool      `db:"setup_completed"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}

func CreateUser(userId, email string, db *sqlx.DB) (*User, error) {
	user := LoadUser(userId, db)
	if user.Id > 0 {
		return user, nil
	}

	query := `INSERT INTO users (
                   user_id,
                   email,
                   setup_completed,
                   created_at,
                   updated_at)
                   VALUES (
                           :user_id,
                           :email,
                           :setup_completed,
                           :created_at,
                           :updated_at)`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"user_id":         userId,
			"email":           email,
			"setup_completed": false,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
		})
	if err != nil {
		return &User{}, err
	}

	return LoadUser(userId, db), nil
}

func LoadUser(userId string, db *sqlx.DB) *User {
	var users []User

	query := `SELECT * FROM users WHERE user_id = $1`
	err := db.Select(&users, query, userId)

	if err != nil {
		return &User{}
	}

	if len(users) > 0 {
		return &users[0]
	}

	return &User{}
}

func LoadUserByEmail(email string, db *sqlx.DB) *User {
	var users []User

	query := `SELECT * FROM users WHERE email = $1`
	err := db.Select(&users, query, email)
	if err != nil {
		return &User{}
	}

	if len(users) > 0 {
		return &users[0]
	}

	return &User{}
}

func (u *User) Save(db *sqlx.DB) error {
	query := `INSERT INTO users (
                   user_id,
                   email,
                   stripe_user_id,
                   stripe_subscription_id,
                   stripe_plan_id,
                   setup_completed)
                   VALUES (
                           :user_id,
                           :email,
                           :stripe_user_id,
                           :stripe_subscription_id,
                           :stripe_plan_id,
                           :setup_completed)
                           ON CONFLICT (user_id, email)
                               DO UPDATE SET
                                             stripe_user_id = :stripe_user_id,
                                             stripe_subscription_id = :stripe_subscription_id,
                                             stripe_plan_id = :stripe_plan_id,
                                             setup_completed = :setup_completed,
                                             updated_at = NOW()`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"user_id":                u.UserId,
			"stripe_user_id":         u.StripeUserId,
			"stripe_subscription_id": u.StripeSubscriptionId,
			"stripe_plan_id":         u.StripePlanId,
			"email":                  u.Email,
			"setup_completed":        u.SetupCompleted,
			"updated_at":             time.Now(),
		})
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Delete(db *sqlx.DB) error {
	query := `DELETE FROM users WHERE id=:id`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"id": u.Id,
		})
	if err != nil {
		return err
	}

	return nil
}
