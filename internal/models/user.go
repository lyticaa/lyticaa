package models

import (
	"database/sql"
	"fmt"
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
	Nickname             sql.NullString `db:"nickname"`
	Picture              sql.NullString `db:"picture"`
	SetupCompleted       bool           `db:"setup_completed"`
	Admin                bool           `db:"admin"`
	Impersonate          *User
	CreatedAt            time.Time `db:"created_at"`
	UpdatedAt            time.Time `db:"updated_at"`
}

var (
	userSortMap = map[int64]string{
		0: "email",
		1: "created_at",
	}
)

func CreateUser(userId, email, nickname, picture string, db *sqlx.DB) (*User, error) {
	user := LoadUser(userId, db)
	if user.Id > 0 {
		return user, nil
	}

	query := `INSERT INTO users (
                   user_id,
                   email,
                   nickname,
                   picture,
                   setup_completed,
                   admin,
                   created_at,
                   updated_at)
                   VALUES (
                           :user_id,
                           :email,
                           :nickname,
                           :picture,
                           :setup_completed,
                           :admin,
                           :created_at,
                           :updated_at)`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"user_id":         userId,
			"email":           email,
			"nickname":        nickname,
			"picture":         picture,
			"setup_completed": false,
			"admin":           false,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
		})
	if err != nil {
		return &User{}, err
	}

	return LoadUser(userId, db), nil
}

func LoadUser(userId string, db *sqlx.DB) *User {
	var user User

	query := `SELECT * FROM users WHERE user_id = $1`
	_ = db.QueryRow(query, userId).Scan(
		&user.Id,
		&user.UserId,
		&user.Email,
		&user.StripeUserId,
		&user.StripeSubscriptionId,
		&user.StripePlanId,
		&user.Nickname,
		&user.Picture,
		&user.SetupCompleted,
		&user.Admin,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return &user
}

func LoadUsers(filter *Filter, db *sqlx.DB) *[]User {
	var users []User

	query := `SELECT * FROM users WHERE created_at BETWEEN $1 AND $2 ORDER BY $3 LIMIT $4 OFFSET $5`
	_ = db.Select(
		&users,
		query,
		filter.StartDate,
		filter.EndDate,
		fmt.Sprintf("%v %v", sortColumn(userSortMap, filter.Sort), filter.Dir),
		filter.Length,
		filter.Start,
	)

	return &users
}

func TotalUsers(filter *Filter, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM users WHERE created_at BETWEEN $1 AND $2`
	_ = db.QueryRow(query, filter.StartDate, filter.EndDate).Scan(&count)

	return count
}

func (u *User) Save(db *sqlx.DB) error {
	query := `UPDATE users SET user_id = :user_id,
                 email = :email,
                 stripe_user_id = :stripe_user_id,
                 stripe_subscription_id = :stripe_subscription_id,
                 stripe_plan_id = :stripe_plan_id,
                 nickname = :nickname,
                 picture = :picture,
                 setup_completed = :setup_completed,
                 updated_at = :updated_at WHERE id = :id`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"user_id":                u.UserId,
			"email":                  u.Email,
			"stripe_user_id":         u.StripeUserId,
			"stripe_subscription_id": u.StripeSubscriptionId,
			"stripe_plan_id":         u.StripePlanId,
			"nickname":               u.Nickname,
			"picture":                u.Picture,
			"setup_completed":        u.SetupCompleted,
			"updated_at":             time.Now(),
			"id":                     u.Id,
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
