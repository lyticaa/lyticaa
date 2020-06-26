package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type CostOfGood struct {
	Id          int64     `db:"id"`
	UserId      int64     `db:"user_id"`
	SKU         string    `db:"sku"`
	Description string    `db:"description"`
	Cost        float64   `db:"cost"`
	StartAt     time.Time `db:"start_at"`
	EndAt       time.Time `db:"end_at"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (c *CostOfGood) Load(db *sqlx.DB) *[]CostOfGood {
	var costOfGoods []CostOfGood

	query := `SELECT cost, start_at, end_at FROM cost_of_goods WHERE user_id = $1 AND sku = $2`
	_ = db.Select(&costOfGoods, query, c.UserId, c.SKU)

	return &costOfGoods
}

func (c *CostOfGood) Save(db *sqlx.DB) error {
	query := `INSERT INTO cost_of_goods (
                           user_id,
                           sku,
                           description,
                           cost,
                           start_at,
                           end_at,
                           created_at,
                           updated_at)
                           VALUES (
                                   :user_id,
                                   :sku,
                                   :description,
                                   :cost,
                                   :start_at,
                                   :end_at,
                                   :created_at,
                                   :updated_at)
                                   ON CONFLICT (user_id, marketplace_id, sku, start_at, end_at)
                                       DO UPDATE SET cost = :cost,
                                                     description = :description,
                                                     updated_at = NOW()`

	_, err := db.NamedExec(query,
		map[string]interface{}{
			"user_id":     c.UserId,
			"sku":         c.SKU,
			"description": c.Description,
			"cost":        c.Cost,
			"start_at":    c.StartAt,
			"end_at":      c.EndAt,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
		})

	if err != nil {
		return err
	}

	return nil
}
