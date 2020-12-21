package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

var (
	userSortMap = map[int64]string{
		0: "email",
		1: "created_at",
	}
)

type UserModel struct {
	ID          int64          `db:"id"`
	UserID      string         `db:"user_id"`
	Email       string         `db:"email"`
	Nickname    sql.NullString `db:"nickname"`
	AvatarURL   sql.NullString `db:"avatar_url"`
	Admin       bool           `db:"admin"`
	Impersonate *UserModel
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (um *UserModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	var user UserModel

	query := `SELECT * FROM users WHERE user_id = $1`
	_ = db.QueryRowxContext(ctx, query, um.UserID).StructScan(&user)

	return user
}

func (um *UserModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} {
	var user UserModel

	query := `SELECT * FROM users WHERE email = $1`
	_ = db.QueryRowxContext(ctx, query, um.Email).StructScan(&user)

	return user
}

func (um *UserModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var users []UserModel

	query := `SELECT * FROM users WHERE created_at BETWEEN $1 AND $2 ORDER BY $3 LIMIT $4 OFFSET $5`
	_ = db.SelectContext(
		ctx,
		&users,
		query,
		filter.StartDate,
		filter.EndDate,
		OrderBy(userSortMap, filter),
		filter.Length,
		filter.Start,
	)

	return users
}

func (um *UserModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM users`
	_ = db.QueryRowContext(ctx, query).Scan(&count)

	return count
}

func (um *UserModel) Create(ctx context.Context, db *sqlx.DB) error {
	query := `INSERT INTO users (
                  user_id,
                  email,
                  nickname,
                  avatar_url,
                  admin,
                  created_at,
                  updated_at)
                  VALUES (
                          :user_id,
                          :email,
                          :nickname,
                          :avatar_url,
                          :admin,
                          :created_at,
                          :updated_at)`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"user_id":    um.UserID,
			"email":      um.Email,
			"nickname":   um.Nickname,
			"avatar_url": um.AvatarURL,
			"admin":      false,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		})
	if err != nil {
		return err
	}

	return nil
}

func (um *UserModel) Update(ctx context.Context, db *sqlx.DB) error {
	query := `UPDATE users SET user_id = :user_id,
                 email = :email,
                 nickname = :nickname,
                 avatar_url = :avatar_url,
                 updated_at = :updated_at WHERE id = :id`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"user_id":    um.UserID,
			"email":      um.Email,
			"nickname":   um.Nickname,
			"avatar_url": um.AvatarURL,
			"updated_at": time.Now(),
			"id":         um.ID,
		})
	if err != nil {
		return err
	}

	return nil
}

func (um *UserModel) Delete(ctx context.Context, db *sqlx.DB) error {
	query := `DELETE FROM users WHERE id = :id`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"id": um.ID,
		})
	if err != nil {
		return err
	}

	return nil
}
