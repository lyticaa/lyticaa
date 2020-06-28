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

func CreateCostOfGood(userId, sku, description string, cost float64, startAt, endAt time.Time, db *sqlx.DB) error {
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
                                   :updated_at)`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"user_id":     userId,
			"sku":         sku,
			"description": description,
			"cost":        cost,
			"start_at":    startAt,
			"end_at":      endAt,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
		})

	if err != nil {
		return err
	}

	return nil
}

func LoadCostOfGood(userId int64, sku string, db *sqlx.DB) *[]CostOfGood {
	var costOfGoods []CostOfGood

	query := `SELECT cost, start_at, end_at FROM cost_of_goods WHERE user_id = $1 AND sku = $2`
	_ = db.Select(&costOfGoods, query, userId, sku)

	return &costOfGoods
}

func (c *CostOfGood) Save(db *sqlx.DB) error {
	query := `UPDATE cost_of_goods SET,
                       sku = :sku,
                       description = :description,
                       cost = :cost,
                       start_at = :start_at,
                       end_at = :end_at,
                       updated_at = :updated_at WHERE user_id = :user_id`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"sku":         c.SKU,
			"description": c.Description,
			"cost":        c.Cost,
			"start_at":    c.StartAt,
			"end_at":      c.EndAt,
			"updated_at":  time.Now(),
			"user_id":     c.UserId,
		})
	if err != nil {
		return err
	}

	return nil
}
