package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type ExchangeRateModel struct {
	ID             int64     `db:"id"`
	ExchangeRateID string    `db:"exchange_rate_id"`
	Code           string    `db:"code"`
	Rate           float64   `db:"rate"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func (er *ExchangeRateModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} { return nil }
func (er *ExchangeRateModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{}  { return nil }

func (er *ExchangeRateModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var exchangeRates []ExchangeRateModel

	query := `SELECT id, exchange_rate_id, code, rate FROM exchange_rates ORDER BY id DESC`
	_ = db.SelectContext(
		ctx,
		&exchangeRates,
		query,
	)

	return exchangeRates
}

func (er *ExchangeRateModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	return int64(0)
}
func (er *ExchangeRateModel) Create(ctx context.Context, db *sqlx.DB) error { return nil }
func (er *ExchangeRateModel) Update(ctx context.Context, db *sqlx.DB) error { return nil }
func (er *ExchangeRateModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
