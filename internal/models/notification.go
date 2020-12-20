package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type NotificationModel struct {
	ID           int64     `db:"id"`
	UserID       int64     `db:"user_id"`
	Notification string    `db:"notification"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

var (
	notificationSortMap = map[int64]string{
		0: "notification",
		1: "created_at",
	}
)

func (nm *NotificationModel) FetchOne(ctx context.Context, db *sqlx.DB) {}
func (nm *NotificationModel) FetchBy(ctx context.Context, db *sqlx.DB)  {}

func (nm *NotificationModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var notifications []NotificationModel

	query := `SELECT notification, created_at FROM notifications WHERE user_id = $1 AND created_at BETWEEN $2 AND $3 ORDER BY $4 LIMIT $5 OFFSET $6`
	_ = db.SelectContext(
		ctx,
		&notifications,
		query,
		nm.UserID,
		filter.StartDate,
		filter.EndDate,
		OrderBy(notificationSortMap, filter),
		filter.Length,
		filter.Start,
	)

	return notifications
}

func (nm *NotificationModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM notifications WHERE user_id = $1`
	_ = db.QueryRowContext(ctx, query, data["UserID"].(string)).Scan(&count)

	return count
}

func (nm *NotificationModel) Create(ctx context.Context, db *sqlx.DB) error { return nil }

func (nm *NotificationModel) Update(ctx context.Context, db *sqlx.DB) error {
	query := `UPDATE notifications SET notification = :notification, updated_at = :updated_at WHERE id = :id`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"user_id":      nm.UserID,
			"notification": nm.Notification,
			"created_at":   time.Now(),
			"updated_at":   time.Now(),
			"id":           nm.ID,
		})

	if err != nil {
		return err
	}

	return nil
}

func (nm *NotificationModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
