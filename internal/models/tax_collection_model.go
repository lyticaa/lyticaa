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

func (t *TaxCollectionModel) Load(db *sqlx.DB) *[]TaxCollectionModel {
	var collectionModels []TaxCollectionModel

	query := `SELECT id,name,created_at,updated_at FROM tax_collection_models ORDER BY id DESC`
	_ = db.Select(&collectionModels, query)
	return &collectionModels
}
