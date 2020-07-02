package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type ExpensesOther struct {
	Id             int64     `db:"id"`
	ExpenseId      string    `db:"expense_id"`
	UserId         string    `db:"user_id"`
	CurrencyId     int64     `db:"currency_id"`
	Description    string    `db:"description"`
	Amount         float64   `db:"amount"`
	CurrencyCode   string    `db:"currency_code"`
	CurrencySymbol string    `db:"currency_symbol"`
	DateTime       time.Time `db:"date_time"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

var (
	expensesOtherSortMap = map[int64]string{
		0: "e.description",
		1: "e.date_time",
		2: "e.amount",
		3: "c.currency_code",
	}
)

func CreateExpensesOther(userId string, currencyId int64, description string, amount float64, dateTime time.Time, db *sqlx.DB) error {
	query := `INSERT INTO expenses_other (
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
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"user_id":     userId,
			"currency_id": currencyId,
			"description": description,
			"amount":      amount,
			"date_time":   dateTime,
			"created_at":  time.Now(),
			"updated_at":  time.Now(),
		})

	if err != nil {
		return err
	}

	return nil
}

func LoadExpensesOther(userId string, filter *Filter, db *sqlx.DB) *[]ExpensesOther {
	var other []ExpensesOther

	query := `SELECT e.expense_id,
       c.code AS currency_code,
       c.symbol AS currency_symbol,
       e.description,
       e.amount,
       e.date_time FROM expenses_other AS e
           LEFT JOIN currencies c ON e.currency_id = c.id WHERE e.user_id = $1 ORDER BY $2 LIMIT $3 OFFSET $4`
	_ = db.Select(
		&other,
		query,
		userId,
		fmt.Sprintf("%v %v", sortColumn(expensesOtherSortMap, filter.Sort), filter.Dir),
		filter.Length,
		filter.Start,
	)

	return &other
}

func TotalExpensesOther(userId string, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM expenses_other WHERE user_id = $1`
	_ = db.QueryRow(query, userId).Scan(&count)

	return count
}

func (e *ExpensesOther) Save(db *sqlx.DB) error {
	query := `UPDATE expenses_other SET
                          currency_id = :currency_id,
                          description = :description,
                          amount = :amount,
                          date_time = :date_time,
                          updated_at = :updated_at WHERE user_id = :user_id
                                                     AND expense_id = :expense_id`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"currency_id": e.CurrencyId,
			"description": e.Description,
			"amount":      e.Amount,
			"date_time":   e.DateTime,
			"updated_at":  time.Now(),
			"user_id":     e.UserId,
			"expense_id":  e.ExpenseId,
		})
	if err != nil {
		return err
	}

	return nil
}
