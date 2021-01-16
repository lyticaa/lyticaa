package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type ExpensesCostOfGoodModel struct {
	ID          int64     `db:"id"`
	ExpenseID   string    `db:"expense_id"`
	ProductID   int64     `db:"product_id"`
	SKU         string    `db:"sku"`
	Marketplace string    `db:"marketplace"`
	Description string    `db:"description"`
	Amount      float64   `db:"amount"`
	FromDate    time.Time `db:"from_date"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

var (
	expensesCostOfGoodsSortMap = map[int64]string{
		0: "p.marketplace",
		1: "p.sku",
		2: "e.description",
		3: "e.from_date",
		4: "e.amount",
	}
)

func (ec *ExpensesCostOfGoodModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	var expensesCostOfGood ExpensesCostOfGoodModel

	query := `SELECT e.id,
       e.expense_id,
       p.product_id,
       e.description,
       e.amount,
       e.from_date,
       e.created_at,
       e.updated_at FROM expenses_cost_of_goods AS e LEFT JOIN products p ON e.product_id = p.id WHERE expense_id = $1`
	_ = db.QueryRowxContext(ctx, query, ec.ExpenseID).StructScan(&expensesCostOfGood)

	return expensesCostOfGood
}

func (ec *ExpensesCostOfGoodModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} { return nil }

func (ec *ExpensesCostOfGoodModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var expensesCostOfGoods []ExpensesCostOfGoodModel

	query := `SELECT e.id, 
       e.expense_id,
       p.product_id,
       p.marketplace,
       p.sku,
       e.description,
       e.amount,
       e.from_date FROM expenses_cost_of_goods AS e
           LEFT JOIN products p ON e.product_id = p.id WHERE p.user_id = $1 ORDER BY $2 LIMIT $3 OFFSET $4`
	_ = db.SelectContext(
		ctx,
		&expensesCostOfGoods,
		query,
		data["UserID"].(string),
		OrderBy(expensesCostOfGoodsSortMap, filter),
		filter.Length,
		filter.Start,
	)

	return expensesCostOfGoods
}

func (ec *ExpensesCostOfGoodModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(e.id) FROM expenses_cost_of_goods AS e LEFT JOIN products p ON e.product_id = p.id WHERE p.user_id = $1`
	_ = db.QueryRowContext(ctx, query, data["UserID"].(string)).Scan(&count)

	return count
}

func (ec *ExpensesCostOfGoodModel) Create(ctx context.Context, db *sqlx.DB) error {
	query := `INSERT INTO expenses_cost_of_goods (
                                    product_id,
                                    description,
                                    amount,
                                    from_date,
                                    created_at,
                                    updated_at)
                                    VALUES (
                                            :product_id,
                                            :description,
                                            :amount,
                                            :from_date,
                                            :created_at,
                                            :updated_at)`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"product_id":  ec.ProductID,
			"description": ec.Description,
			"amount":      ec.Amount,
			"from_date":   ec.FromDate,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
		})
	if err != nil {
		return err
	}

	return nil
}

func (ec *ExpensesCostOfGoodModel) Update(ctx context.Context, db *sqlx.DB) error {
	query := `UPDATE expenses_cost_of_goods SET
                                  description = :description,
                                  amount = :amount,
                                  from_date = :from_date,
                                  updated_at = :updated_at WHERE expense_id = :expense_id
                                                             AND product_id = :product_id`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"description": ec.Description,
			"amount":      ec.Amount,
			"from_date":   ec.FromDate,
			"updated_at":  time.Now(),
			"expense_id":  ec.ExpenseID,
			"product_id":  ec.ProductID,
		})
	if err != nil {
		return err
	}

	return nil
}

func (ec *ExpensesCostOfGoodModel) Delete(ctx context.Context, db *sqlx.DB) error {
	query := `DELETE FROM expenses_cost_of_goods WHERE id = :id AND expense_id = :expense_id`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"id":         ec.ID,
			"expense_id": ec.ExpenseID,
		})
	if err != nil {
		return err
	}

	return nil
}
