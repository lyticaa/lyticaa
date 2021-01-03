package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type CurrencyModel struct {
	ID         int64     `db:"id"`
	CurrencyID string    `db:"currency_id"`
	Code       string    `db:"code"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (cm *CurrencyModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	var currency CurrencyModel

	query := `SELECT id, currency_id, code, created_at, updated_at FROM currencies WHERE currency_id = $1`
	_ = db.QueryRowxContext(ctx, query, cm.CurrencyID).StructScan(&currency)

	return currency
}

func (cm *CurrencyModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} { return nil }

func (cm *CurrencyModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var currency []CurrencyModel

	query := `SELECT id, currency_id, code, FROM currencies`
	_ = db.SelectContext(ctx, &currency, query)

	return currency
}

func (cm *CurrencyModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM currencies`
	_ = db.QueryRowContext(ctx, query).Scan(&count)

	return count
}

func (cm *CurrencyModel) Create(ctx context.Context, db *sqlx.DB) error { return nil }
func (cm *CurrencyModel) Update(ctx context.Context, db *sqlx.DB) error { return nil }
func (cm *CurrencyModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
