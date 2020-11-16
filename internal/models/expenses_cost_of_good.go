package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type ExpensesCostOfGood struct {
	ID          int64     `db:"id"`
	ExpenseID   string    `db:"expense_id"`
	ProductID   string    `db:"product_id"`
	SKU         string    `db:"sku"`
	Marketplace string    `db:"marketplace"`
	Description string    `db:"description"`
	Amount      float64   `db:"amount"`
	FromDate    time.Time `db:"from_date"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type ExpenseCostOfGoodSendData struct {
	ExpenseID   string    `json:"expenseID"`
	UserID      string    `json:"userID"`
	SKU         string    `json:"sku"`
	Marketplace string    `json:"marketplace"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	DateTime    time.Time `json:"dateTime"`
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

func CreateExpensesCostOfGood(productID int64, description string, amount float64, fromDate time.Time, db *sqlx.DB) (string, error) {
	var expenseID string

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
											:updated_at)
									RETURNING expense_id`
	rows, err := db.NamedQuery(query,
		map[string]interface{}{
			"product_id":  productID,
			"description": description,
			"amount":      amount,
			"from_date":   fromDate,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
		})

	if err != nil {
		return "", err
	}

	if rows.Next() {

		if err := rows.Scan(&expenseID); err != nil {
			return "", err
		}
	}

	if err := rows.Close(); err != nil {
		return "", err
	}

	return expenseID, nil
}

func LoadExpensesCostOfGoods(userID string, filter *Filter, db *sqlx.DB) *[]ExpensesCostOfGood {
	var costOfGoods []ExpensesCostOfGood

	query := `SELECT e.id, 
       e.expense_id,
       p.product_id,
       p.marketplace,
       p.sku,
       e.description,
       e.amount,
       e.from_date FROM expenses_cost_of_goods AS e
           LEFT JOIN products p ON e.product_id = p.id WHERE p.user_id = $1 ORDER BY $2 LIMIT $3 OFFSET $4`
	_ = db.Select(
		&costOfGoods,
		query,
		userID,
		fmt.Sprintf("%v %v", sortColumn(expensesCostOfGoodsSortMap, filter.Sort), filter.Dir),
		filter.Length,
		filter.Start,
	)

	return &costOfGoods
}

func LoadExpensesCostOfGood(expenseID string, db *sqlx.DB) *ExpensesCostOfGood {
	var costOfGood ExpensesCostOfGood

	query := `SELECT e.id,
       e.expense_id,
       p.product_id,
       e.description,
       e.amount,
       e.from_date,
       e.created_at,
       e.updated_at FROM expenses_cost_of_goods AS e LEFT JOIN products p ON e.product_id = p.id WHERE expense_id = $1`
	_ = db.QueryRow(query, expenseID).Scan(
		&costOfGood.ID,
		&costOfGood.ExpenseID,
		&costOfGood.ProductID,
		&costOfGood.Description,
		&costOfGood.Amount,
		&costOfGood.FromDate,
		&costOfGood.CreatedAt,
		&costOfGood.UpdatedAt,
	)

	return &costOfGood
}

func TotalExpensesCostOfGoods(userID string, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(e.id) FROM expenses_cost_of_goods AS e LEFT JOIN products p ON e.product_id = p.id WHERE p.user_id = $1`
	_ = db.QueryRow(query, userID).Scan(&count)

	return count
}

func (e *ExpensesCostOfGood) Save(db *sqlx.DB) error {
	query := `UPDATE expenses_cost_of_goods SET
                                  description = :description,
                                  amount = :amount,
                                  from_date = :from_date,
                                  updated_at = :updated_at WHERE expense_id = :expense_id`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"description": e.Description,
			"amount":      e.Amount,
			"from_date":   e.FromDate,
			"updated_at":  time.Now(),
			"expense_id":  e.ExpenseID,
		})
	if err != nil {
		return err
	}

	return nil
}

func (e *ExpensesCostOfGood) Delete(db *sqlx.DB) error {
	query := `DELETE FROM expenses_cost_of_goods WHERE id = :id
                                     AND expense_id = :expense_id`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"id":         e.ID,
			"expense_id": e.ExpenseID,
		})
	if err != nil {
		return err
	}

	return nil
}
