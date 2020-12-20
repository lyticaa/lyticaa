package data

import (
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/chart"

	"github.com/jmoiron/sqlx"
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
