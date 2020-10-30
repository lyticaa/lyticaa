package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int64     `db:"id"`
	ProductID   string    `db:"product_id"`
	UserID      string    `db:"user_id"`
	SKU         string    `db:"sku"`
	Marketplace string    `db:"marketplace"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func LoadProducts(userID string, db *sqlx.DB) *[]Product {
	var products []Product

	query := `SELECT product_id, sku, marketplace, description, created_at, updated_at FROM products WHERE user_id = $1`
	_ = db.Select(&products, query, userID)

	return &products
}

func LoadProduct(userID, productID string, db *sqlx.DB) *Product {
	var product Product

	query := `SELECT id,
       product_id,
       user_id,
       sku,
       marketplace,
       description,
       created_at,
       updated_at FROM products WHERE user_id = $1 
                                  AND product_id = $2`
	_ = db.QueryRow(query, userID, productID).Scan(
		&product.ID,
		&product.ProductID,
		&product.UserID,
		&product.SKU,
		&product.Marketplace,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	return &product
}

func LoadProductByID(userID string, productID int64, db *sqlx.DB) *Product {
	var product Product

	query := `SELECT id,
       product_id,
       user_id,
       sku,
       marketplace,
       description,
       created_at,
       updated_at FROM products WHERE user_id = $1 
                                  AND id = $2`
	_ = db.QueryRow(query, userID, productID).Scan(
		&product.ID,
		&product.ProductID,
		&product.UserID,
		&product.SKU,
		&product.Marketplace,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	return &product
}

func (p *Product) Save(db *sqlx.DB) error {
	query := `UPDATE products SET
                    sku = :sku,
                    marketplace = :marketplace,
                    description = :description,
                    updated_at = :updated_at WHERE user_id = :user_id
                                               AND product_id = :product_id`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"sku":         p.SKU,
			"marketplace": p.Marketplace,
			"description": p.Description,
			"updated_at":  time.Now(),
			"user_id":     p.UserID,
			"product_id":  p.ProductID,
		})
	if err != nil {
		return err
	}

	return nil
}
