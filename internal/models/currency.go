package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Currency struct {
	Id         int64     `db:"id"`
	CurrencyId string    `db:"currency_id"`
	Code       string    `db:"code"`
	Symbol     string    `db:"symbol"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func LoadCurrencies(db *sqlx.DB) *[]Currency {
	var currency []Currency

	query := `SELECT id, currency_id, code, symbol FROM currencies`
	_ = db.Select(&currency, query)

	return &currency
}

func LoadCurrency(currencyId string, db *sqlx.DB) *Currency {
	var currency Currency

	query := `SELECT id, currency_id, code, symbol, created_at, updated_at FROM currencies WHERE currency_id = $1`
	_ = db.QueryRow(query, currencyId).Scan(
		&currency.Id,
		&currency.CurrencyId,
		&currency.Code,
		&currency.Symbol,
		&currency.CreatedAt,
		&currency.UpdatedAt,
	)

	return &currency
}
