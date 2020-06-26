package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type ExchangeRate struct {
	Id            int64     `db:"id"`
	MarketplaceId int64     `db:"marketplace_id"`
	Code          string    `db:"code"`
	Symbol        string    `db:"symbol"`
	Rate          float64   `db:"rate"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func (e *ExchangeRate) Load(db *sqlx.DB) *[]ExchangeRate {
	var exchangeRates []ExchangeRate

	query := `SELECT id,code,symbol,rate,created_at,updated_at FROM exchange_rates ORDER BY id DESC`
	_ = db.Select(&exchangeRates, query)

	return &exchangeRates
}
