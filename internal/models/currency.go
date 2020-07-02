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

	query := `SELECT currency_id, code, symbol FROM currencies`
	_ = db.Select(&currency, query)

	return &currency
}
