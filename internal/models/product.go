package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	Id          int64     `db:"id"`
	ProductId   string    `db:"product_id"`
	UserId      string    `db:"user_id"`
	SKU         string    `db:"sku"`
	Marketplace string    `db:"marketplace"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func LoadProducts(userId string, db *sqlx.DB) *[]Product {
	var products []Product

	query := `SELECT product_id, sku, marketplace, description, created_at, updated_at FROM products WHERE user_id = $1`
	_ = db.Select(&products, query, userId)

	return &products
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
			"user_id":     p.UserId,
			"product_id":  p.ProductId,
		})
	if err != nil {
		return err
	}

	return nil
}
