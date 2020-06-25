package dashboard

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/amazon"
	"gitlab.com/getlytica/lytica-app/internal/core/app/chart"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Dashboard struct {
	amazon       *amazon.Amazon
	chart        *chart.Chart
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewDashboard(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Dashboard {
	return &Dashboard{
		amazon:       amazon.NewAmazon(db),
		chart:        chart.NewChart(),
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
