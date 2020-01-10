package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Currency struct {
	Id        int64
	Name      string
	Symbol    string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func GetCurrencies(db *sqlx.DB) []Currency {
	var currencies []Currency
	_ = db.Select(&currencies, "SELECT id,name,symbol,created_at,updated_at FROM currencies ORDER BY id DESC")

	return currencies
}
