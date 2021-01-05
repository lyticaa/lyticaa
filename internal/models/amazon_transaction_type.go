package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type AmazonTransactionTypeModel struct {
	ID                      int64     `db:"id"`
	AmazonTransactionTypeID string    `db:"amazon_transaction_type_id"`
	Name                    string    `db:"name"`
	CreatedAt               time.Time `db:"created_at"`
	UpdatedAt               time.Time `db:"updated_at"`
}

func (at *AmazonTransactionTypeModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	return nil
}
func (at *AmazonTransactionTypeModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} {
	return nil
}

func (at *AmazonTransactionTypeModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var amazonTransactionTypes []AmazonTransactionTypeModel

	query := `SELECT id, amazon_transaction_type_id, name FROM amazon_transaction_types ORDER BY id DESC`
	_ = db.SelectContext(
		ctx,
		&amazonTransactionTypes,
		query,
	)

	return &amazonTransactionTypes
}

func (at *AmazonTransactionTypeModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	return int64(0)
}
func (at *AmazonTransactionTypeModel) Create(ctx context.Context, db *sqlx.DB) error { return nil }
func (at *AmazonTransactionTypeModel) Update(ctx context.Context, db *sqlx.DB) error { return nil }
func (at *AmazonTransactionTypeModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
