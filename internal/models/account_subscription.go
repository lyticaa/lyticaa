package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type AccountSubscriptionModel struct {
	ID                    int64          `db:"id"`
	AccountSubscriptionID string         `db:"account_subscription_id"`
	UserID                int64          `db:"user_id"`
	StripeSubscriptionID  sql.NullString `db:"stripe_subscription_id"`
	StripePlanID          sql.NullString `db:"stripe_plan_id"`
	CreatedAt             time.Time      `db:"created_at"`
	UpdatedAt             time.Time      `db:"updated_at"`
}

func (as *AccountSubscriptionModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	var accountSubscription AccountSubscriptionModel

	query := `SELECT * FROM account_subscriptions WHERE user_id = $1`
	_ = db.QueryRowxContext(ctx, query, as.UserID).StructScan(&accountSubscription)

	return accountSubscription
}

func (as *AccountSubscriptionModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} { return nil }
func (as *AccountSubscriptionModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	return nil
}

func (as *AccountSubscriptionModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	return int64(0)
}

func (as *AccountSubscriptionModel) Create(ctx context.Context, db *sqlx.DB) error {
	query := `INSERT INTO account_subscriptions (
                                   user_id,
                                   stripe_subscription_id,
                                   stripe_plan_id,
                                   created_at,
                                   updated_at)
                                   VALUES (
                                           :user_id,
                                           :stripe_subscription_id,
                                           :stripe_plan_id,
                                           :created_at,
                                           :updated_at)`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"user_id":                as.UserID,
			"stripe_subscription_id": as.StripeSubscriptionID,
			"stripe_plan_id":         as.StripePlanID,
			"created_at":             time.Now(),
			"updated_at":             time.Now(),
		})
	if err != nil {
		return err
	}

	return nil
}

func (as *AccountSubscriptionModel) Update(ctx context.Context, db *sqlx.DB) error {
	query := `UPDATE account_subscriptions SET stripe_subscription_id = :stripe_subscription_id,
                               stripe_plan_id = :stripe_plan_id,
                               updated_at = :updated_at WHERE account_subscription_id = :account_subscription_id
                                                          AND user_id = :user_id`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"stripe_subscription_id":  as.StripeSubscriptionID,
			"stripe_plan_id":          as.StripePlanID,
			"updated_at":              time.Now(),
			"account_subscription_id": as.AccountSubscriptionID,
			"user_id":                 as.UserID,
		})

	if err != nil {
		return err
	}

	return nil
}

func (as *AccountSubscriptionModel) Delete(ctx context.Context, db *sqlx.DB) error {
	query := `DELETE FROM account_subscriptions WHERE id = :id
                                    AND account_subscription_id = :account_subscription_id
                                    AND user_id = :user_id`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"id":                      as.ID,
			"account_subscription_id": as.AccountSubscriptionID,
			"user_id":                 as.UserID,
		})
	if err != nil {
		return err
	}

	return nil
}
