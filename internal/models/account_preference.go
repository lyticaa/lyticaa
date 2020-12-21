package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type AccountPreferenceModel struct {
	ID                  int64        `db:"id"`
	AccountPreferenceID string       `db:"account_preference_id"`
	UserID              int64        `db:"user_id"`
	SetupCompleted      bool         `db:"setup_completed"`
	MailingList         sql.NullBool `db:"mailing_list"`
	CreatedAt           time.Time    `db:"created_at"`
	UpdatedAt           time.Time    `db:"updated_at"`
}

func (ap *AccountPreferenceModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} { return nil }
func (ap *AccountPreferenceModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{}  { return nil }

func (ap *AccountPreferenceModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var accountPreferences AccountPreferenceModel

	query := `SELECT * FROM account_preferences WHERE user_id = $1`
	_ = db.QueryRowxContext(ctx, query, ap.UserID).StructScan(&accountPreferences)

	return accountPreferences
}

func (ap *AccountPreferenceModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	return int64(0)
}

func (ap *AccountPreferenceModel) Create(ctx context.Context, db *sqlx.DB) error {
	query := `INSERT INTO account_preferences (
                  user_id,
                  setup_completed,
                  created_at,
                  updated_at)
                  VALUES (
                          :user_id,
                          :setup_completed,
                          :created_at,
                          :updated_at)`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"user_id":         ap.UserID,
			"setup_completed": false,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
		})
	if err != nil {
		return err
	}

	return nil
}

func (ap *AccountPreferenceModel) Update(ctx context.Context, db *sqlx.DB) error {
	query := `UPDATE account_preferences SET setup_completed = :setup_completed,
                               mailing_list = :mailing_list,
                               updated_at = :updated_at WHERE account_preference_id = :account_preference_id
                                                          AND user_id = :user_id`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"setup_completed":       ap.SetupCompleted,
			"mailing_list":          ap.MailingList,
			"updated_at":            time.Now(),
			"account_preference_id": ap.AccountPreferenceID,
			"user_id":               ap.UserID,
		})

	if err != nil {
		return err
	}

	return nil
}

func (ap *AccountPreferenceModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
