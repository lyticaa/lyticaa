package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type AmazonMarketplaceModel struct {
	ID                  int64     `db:"id"`
	AmazonMarketplaceID string    `db:"amazon_marketplace_id"`
	Name                string    `db:"name"`
	ExchangeRateID      int64     `db:"exchange_rate_id"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}

func (am *AmazonMarketplaceModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	var amazonMarketplace AmazonMarketplaceModel

	query := `SELECT id,
       amazon_marketplace_id,
       name,
       exchange_rate_id,
       created_at,
       updated_at FROM amazon_marketplaces where name = $1`
	_ = db.QueryRowxContext(ctx, query, am.Name).StructScan(&amazonMarketplace)

	return amazonMarketplace
}

func (am *AmazonMarketplaceModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} { return nil }

func (am *AmazonMarketplaceModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var amazonMarketplaces []AmazonMarketplaceModel

	query := `SELECT id, amazon_marketplace_id, name, exchange_rate_id FROM amazon_marketplaces ORDER BY id DESC`
	_ = db.SelectContext(
		ctx,
		&amazonMarketplaces,
		query,
	)

	return &amazonMarketplaces
}

func (am *AmazonMarketplaceModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	return int64(0)
}
func (am *AmazonMarketplaceModel) Create(ctx context.Context, db *sqlx.DB) error { return nil }
func (am *AmazonMarketplaceModel) Update(ctx context.Context, db *sqlx.DB) error { return nil }
func (am *AmazonMarketplaceModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
