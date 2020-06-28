package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type TransactionType struct {
	Id        int64
	Name      string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func LoadTransactionTypes(db *sqlx.DB) *[]TransactionType {
	var txnTypes []TransactionType

	query := `SELECT id, name, created_at, updated_at FROM transaction_types ORDER BY id DESC`
	_ = db.Select(&txnTypes, query)

	return &txnTypes
}
