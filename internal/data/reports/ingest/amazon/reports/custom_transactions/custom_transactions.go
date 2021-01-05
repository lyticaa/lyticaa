package custom_transactions

import (
	"github.com/jmoiron/sqlx"
)

type CustomTransactions struct {
	db *sqlx.DB
}

func NewCustomTransactions(db *sqlx.DB) *CustomTransactions {
	return &CustomTransactions{
		db: db,
	}
}
