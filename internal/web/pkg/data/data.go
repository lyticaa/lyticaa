package data

import (
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/chart"

	"github.com/jmoiron/sqlx"
)

const (
	expensesCostOfGoods = "expenses_cost_of_goods"
	expensesOther       = "expenses_other"
)

type Data struct {
	chart *chart.Chart
	db    *sqlx.DB
}

func NewData(db *sqlx.DB) *Data {
	return &Data{
		chart: chart.NewChart(),
		db:    db,
	}
}
