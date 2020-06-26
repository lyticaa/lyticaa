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

func (u *User) Load(db *sqlx.DB) {
	query := `SELECT * FROM users WHERE user_id = $1`
	_ = db.QueryRow(query, u.UserId).Scan(
		&u.Id,
		&u.UserId,
		&u.Email,
		&u.StripeUserId,
		&u.StripeSubscriptionId,
		&u.StripePlanId,
		&u.SetupCompleted,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
}

func (u *User) LoadByEmail(db *sqlx.DB) {
	query := `SELECT * FROM users WHERE email = $1`
	_ = db.QueryRow(query, u.Email).Scan(
		&u.Id,
		&u.UserId,
		&u.Email,
		&u.StripeUserId,
		&u.StripeSubscriptionId,
		&u.StripePlanId,
		&u.SetupCompleted,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
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
