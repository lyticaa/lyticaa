package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type AmazonFulfillmentModel struct {
	ID                  int64     `db:"id"`
	AmazonFulfillmentID string    `db:"amazon_fulfillment_id"`
	Name                string    `db:"name"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}

func (af *AmazonFulfillmentModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	var amazonFulfillment AmazonFulfillmentModel

	query := `SELECT * FROM amazon_fulfillments WHERE name = $1`
	_ = db.QueryRowxContext(ctx, query, af.Name).StructScan(&amazonFulfillment)

	return amazonFulfillment
}

func (af *AmazonFulfillmentModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{}  { return nil }

func (af *AmazonFulfillmentModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var amazonFulfillments []AmazonFulfillmentModel

	query := `SELECT id, amazon_fulfillment_id, name FROM amazon_fulfillments ORDER BY id DESC`
	_ = db.SelectContext(
		ctx,
		&amazonFulfillments,
		query,
	)

	return amazonFulfillments
}

func (af *AmazonFulfillmentModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	return int64(0)
}

func (af *AmazonFulfillmentModel) Create(ctx context.Context, db *sqlx.DB) error {
	query := `INSERT INTO amazon_fulfillments (
                                 name,
                                 created_at,
                                 updated_at)
                                 VALUES (
                                         :name,
                                         :created_at,
                                         :updated_at)`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"name":         af.Name,
			"created_at":      time.Now(),
			"updated_at":      time.Now(),
		})
	if err != nil {
		return err
	}

	return nil
}

func (af *AmazonFulfillmentModel) Update(ctx context.Context, db *sqlx.DB) error { return nil }
func (af *AmazonFulfillmentModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
