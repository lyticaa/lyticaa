package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type ExpensesCostOfGood struct {
	Id          int64     `db:"id"`
	UserId      string    `db:"user_id"`
	Marketplace string    `db:"marketplace"`
	SKU         string    `db:"sku"`
	Description string    `db:"description"`
	Cost        float64   `db:"cost"`
	FromDate    time.Time `db:"from_date"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

var (
	expensesCostOfGoodsSortMap = map[int64]string{
		0: "marketplace",
		1: "sku",
		2: "description",
		3: "cost",
		4: "from_date",
	}
)

func CreateExpensesCostOfGood(userId, marketplace, sku, description string, cost float64, fromDate time.Time, db *sqlx.DB) error {
	query := `INSERT INTO expenses_cost_of_goods (
                                    user_id,
                                    marketplace,
                                    sku,
                                    description,
                                    cost,
                                    from_date,
                                    created_at,
                                    updated_at)
                                    VALUES (
                                            :user_id,
                                            :marketplace,
                                            :sku,
                                            :description,
                                            :cost,
                                            :from_date,
                                            :created_at,
                                            :updated_at)`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"user_id":     userId,
			"marketplace": marketplace,
			"sku":         sku,
			"description": description,
			"cost":        cost,
			"from_date":   fromDate,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
		})

	if err != nil {
		return err
	}

	return nil
}

func LoadExpensesCostOfGoods(userId string, filter *Filter, db *sqlx.DB) *[]ExpensesCostOfGood {
	var costOfGoods []ExpensesCostOfGood

	query := `SELECT cost, from_date FROM expenses_cost_of_goods WHERE user_id = $1 ORDER BY $2 LIMIT $3 OFFSET $4`
	_ = db.Select(
		&costOfGoods,
		query,
		userId,
		fmt.Sprintf("%v %v", sortColumn(expensesCostOfGoodsSortMap, filter.Sort), filter.Dir),
		filter.Length,
		filter.Start,
	)

	return &costOfGoods
}

func TotalExpensesCostOfGoods(userId string, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM expenses_cost_of_goods WHERE user_id = $1`
	_ = db.QueryRow(query, userId).Scan(&count)

	return count
}

func (e *ExpensesCostOfGood) Save(db *sqlx.DB) error {
	query := `UPDATE expenses_cost_of_goods SET
                                  sku = :sku,
                                  description = :description,
                                  cost = :cost,
                                  from_date = :from_date,
                                  updated_at = :updated_at WHERE user_id = :user_id`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"sku":         e.SKU,
			"description": e.Description,
			"cost":        e.Cost,
			"from_date":   e.FromDate,
			"updated_at":  time.Now(),
			"user_id":     e.UserId,
		})
	if err != nil {
		return err
	}

	return nil
}
