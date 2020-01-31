package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type TaxCollectionModel struct {
	Id        int64
	Name      string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func GetTaxCollectionModels(db *sqlx.DB) []TaxCollectionModel {
	var collections []TaxCollectionModel

	err := db.Select(&collections, "SELECT id,name,created_at,updated_at FROM tax_collection_models ORDER BY id DESC")
	if err != nil {
		panic(err)
	}

	return collections
}