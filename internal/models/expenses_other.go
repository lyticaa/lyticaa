package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type ExpensesOther struct {
	Id          int64     `db:"id"`
	UserId      string    `db:"user_id"`
	Description string    `db:"description"`
	Cost        float64   `db:"cost"`
	DateTime    time.Time `db:"date_time"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

var (
	expensesOtherSortMap = map[int64]string{
		0: "description",
		1: "date_time",
		2: "cost",
	}
)

func CreateExpensesOther(userId, description string, cost float64, dateTime time.Time, db *sqlx.DB) error {
	query := `INSERT INTO expenses_other (
                            user_id,
                            description,
                            cost,
                            date_time,
                            created_at,
                            updated_at)
                            VALUES (
                                    :user_id,
                                    :description,
                                    :cost,
                                    :date_time,
                                    :created_at,
                                    :updated_at)`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"user_id":     userId,
			"description": description,
			"cost":        cost,
			"from_date":   dateTime,
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

	query := `SELECT cost, from_date FROM expenses_cost_of_goods WHERE user_id = $1 ORDER BY $2 LIMIT $3 OFFSET $4`
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
                          description = :description,
                          cost = :cost,
                          date_time = :date_time,
                          updated_at = :updated_at WHERE user_id = :user_id`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"description": e.Description,
			"cost":        e.Cost,
			"date_time":   e.DateTime,
			"updated_at":  time.Now(),
			"user_id":     e.UserId,
		})
	if err != nil {
		return err
	}

	return nil
}
