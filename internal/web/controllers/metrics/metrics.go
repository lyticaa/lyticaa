package metrics

import (
	"gitlab.com/getlytica/lytica-app/internal/web/lib/chart"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Metrics struct {
	chart        *chart.Chart
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewMetrics(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Metrics {
	return &Metrics{
		chart:        chart.NewChart(),
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
