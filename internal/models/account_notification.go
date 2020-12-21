package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type AccountNotificationModel struct {
	ID                    int64     `db:"id"`
	AccountNotificationID string    `db:"account_notification_id"`
	UserID                int64     `db:"user_id"`
	Notification          string    `db:"notification"`
	CreatedAt             time.Time `db:"created_at"`
	UpdatedAt             time.Time `db:"updated_at"`
}

var (
	accountNotificationSortMap = map[int64]string{
		0: "notification",
		1: "created_at",
	}
)

func (an *AccountNotificationModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	return nil
}
func (an *AccountNotificationModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} { return nil }

func (an *AccountNotificationModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var accountNotifications []AccountNotificationModel

	query := `SELECT
       account_notification_id,
       notification,
       created_at FROM account_notifications WHERE user_id = $1
                                               AND created_at BETWEEN $2 AND $3 ORDER BY $4 LIMIT $5 OFFSET $6`
	_ = db.SelectContext(
		ctx,
		&accountNotifications,
		query,
		an.UserID,
		filter.StartDate,
		filter.EndDate,
		OrderBy(accountNotificationSortMap, filter),
		filter.Length,
		filter.Start,
	)

	return accountNotifications
}

func (an *AccountNotificationModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM account_notifications WHERE user_id = $1`
	_ = db.QueryRowContext(ctx, query, data["UserID"].(string)).Scan(&count)

	return count
}

func (an *AccountNotificationModel) Create(ctx context.Context, db *sqlx.DB) error { return nil }
func (an *AccountNotificationModel) Update(ctx context.Context, db *sqlx.DB) error { return nil }
func (an *AccountNotificationModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
