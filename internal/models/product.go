package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type ProductModel struct {
	ID          int64     `db:"id"`
	ProductID   string    `db:"product_id"`
	UserID      int64     `db:"user_id"`
	SKU         string    `db:"sku"`
	Marketplace string    `db:"marketplace"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (pm *ProductModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	var product ProductModel

	query := `SELECT id,
       product_id,
       user_id,
       sku,
       marketplace,
       description,
       created_at,
       updated_at FROM products WHERE user_id = $1 
                                  AND product_id = $2`
	_ = db.QueryRowxContext(ctx, query, pm.UserID, pm.ProductID).StructScan(&product)

	return product
}

func (pm *ProductModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} {
	var product ProductModel

	query := `SELECT id,
       product_id,
       user_id,
       sku,
       marketplace,
       description,
       created_at,
       updated_at FROM products WHERE user_id = $1 
                                  AND id = $2`
	_ = db.QueryRowxContext(ctx, query, pm.UserID, pm.ProductID).StructScan(&product)

	return product
}

func (pm *ProductModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var products []ProductModel

	query := `SELECT product_id, sku, marketplace, description, created_at, updated_at FROM products WHERE user_id = $1`
	_ = db.SelectContext(ctx, &products, query, pm.UserID)

	return &products
}

func (pm *ProductModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	return int64(0)
}
func (pm *ProductModel) Create(ctx context.Context, db *sqlx.DB) error { return nil }

func (pm *ProductModel) Update(ctx context.Context, db *sqlx.DB) error {
	query := `UPDATE products SET
                    sku = :sku,
                    marketplace = :marketplace,
                    description = :description,
                    updated_at = :updated_at WHERE product_id = :product_id
                                               AND user_id = :user_id`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"sku":         pm.SKU,
			"marketplace": pm.Marketplace,
			"description": pm.Description,
			"updated_at":  time.Now(),
			"user_id":     pm.UserID,
			"product_id":  pm.ProductID,
		})
	if err != nil {
		return err
	}

	return nil
}

func (pm *ProductModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
