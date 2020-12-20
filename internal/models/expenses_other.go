package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type ExpensesOtherModel struct {
	ID           int64     `db:"id"`
	ExpenseID    string    `db:"expense_id"`
	UserID       int64     `db:"user_id"`
	CurrencyID   int64     `db:"currency_id"`
	Description  string    `db:"description"`
	Amount       float64   `db:"amount"`
	CurrencyCode string    `db:"currency_code"`
	DateTime     time.Time `db:"date_time"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

var (
	expensesOtherSortMap = map[int64]string{
		0: "e.description",
		1: "e.date_time",
		2: "e.amount",
		3: "c.currency_code",
	}
)

func (eo *ExpensesOtherModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	var expensesOther ExpensesOtherModel

	query := `SELECT e.id,
       e.expense_id,
       e.user_id,
       e.currency_id,
       c.code AS currency_code,
       c.symbol AS currency_symbol,
       e.description,
       e.amount,
       e.date_time,
       e.created_at,
       e.updated_at FROM expenses_others AS e LEFT JOIN currencies c ON e.currency_id = c.id WHERE e.expense_id = $1`
	_ = db.QueryRowxContext(ctx, query, eo.ExpenseID).StructScan(&expensesOther)

	return expensesOther
}

func (eo *ExpensesOtherModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} { return nil }

func (eo *ExpensesOtherModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	var expensesOthers []ExpensesOtherModel

	query := `SELECT e.expense_id,
       e.currency_id,
       c.code AS currency_code,
       e.description,
       e.amount,
       e.date_time FROM expenses_others AS e
           LEFT JOIN currencies c ON e.currency_id = c.id WHERE e.user_id = $1 ORDER BY $2 LIMIT $3 OFFSET $4`
	_ = db.SelectContext(
		ctx,
		&expensesOthers,
		query,
		data["UserID"].(string),
		OrderBy(expensesOtherSortMap, filter),
		filter.Length,
		filter.Start,
	)

	return expensesOthers
}

func (eo *ExpensesOtherModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM expenses_others WHERE user_id = $1`
	_ = db.QueryRow(query, data["UserID"].(string)).Scan(&count)

	return count
}

func (eo *ExpensesOtherModel) Create(ctx context.Context, db *sqlx.DB) error {
	query := `INSERT INTO expenses_others (
                            user_id,
                            currency_id,
                            description,
                            amount,
                            date_time,
                            created_at,
                            updated_at)
                            VALUES (
                                    :user_id,
                                    :currency_id,
                                    :description,
                                    :amount,
                                    :date_time,
                                    :created_at,
                                    :updated_at)`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"user_id":     eo.UserID,
			"currency_id": eo.CurrencyID,
			"description": eo.Description,
			"amount":      eo.Amount,
			"date_time":   eo.DateTime,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
		})
	if err != nil {
		return err
	}

	return nil
}

func (eo *ExpensesOtherModel) Update(ctx context.Context, db *sqlx.DB) error {
	query := `UPDATE expenses_others SET
                          currency_id = :currency_id,
                          description = :description,
                          amount = :amount,
                          date_time = :date_time,
                          updated_at = :updated_at WHERE expense_id = :expense_id`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"currency_id": eo.CurrencyID,
			"description": eo.Description,
			"amount":      eo.Amount,
			"date_time":   eo.DateTime,
			"updated_at":  time.Now(),
			"expense_id":  eo.ExpenseID,
		})
	if err != nil {
		return err
	}

	return nil
}

func (eo *ExpensesOtherModel) Delete(ctx context.Context, db *sqlx.DB) error {
	query := `DELETE FROM expenses_others WHERE id = :id AND expense_id = :expense_id`
	_, err := db.NamedExecContext(ctx, query,
		map[string]interface{}{
			"id":         eo.ID,
			"expense_id": eo.ExpenseID,
		})
	if err != nil {
		return err
	}

	return nil
}
