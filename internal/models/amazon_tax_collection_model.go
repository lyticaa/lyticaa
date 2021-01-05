package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type AmazonTaxCollectionModelModel struct {
	ID                         int64     `db:"id"`
	AmazonTaxCollectionModelID string    `db:"amazon_tax_collection_model_id"`
	Name                       string    `db:"name"`
	CreatedAt                  time.Time `db:"created_at"`
	UpdatedAt                  time.Time `db:"updated_at"`
}

func (at *AmazonTaxCollectionModelModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	return nil
}
func (at *AmazonTaxCollectionModelModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} {
	return nil
}

func (at *AmazonTaxCollectionModelModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var amazonTaxCollectionModels []AmazonTaxCollectionModelModel

	query := `SELECT id, amazon_tax_collection_model_id, name FROM amazon_tax_collection_models ORDER BY id DESC`
	_ = db.SelectContext(
		ctx,
		&amazonTaxCollectionModels,
		query,
	)

	return &amazonTaxCollectionModels
}

func (at *AmazonTaxCollectionModelModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	return int64(0)
}
func (at *AmazonTaxCollectionModelModel) Create(ctx context.Context, db *sqlx.DB) error { return nil }
func (at *AmazonTaxCollectionModelModel) Update(ctx context.Context, db *sqlx.DB) error { return nil }
func (at *AmazonTaxCollectionModelModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
