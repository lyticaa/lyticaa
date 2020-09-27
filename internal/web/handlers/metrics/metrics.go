package metrics

import (
	"gitlab.com/lyticaa/lyticaa-app/internal/web/pkg/data"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Metrics struct {
	data         *data.Data
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewMetrics(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Metrics {
	return &Metrics{
		data:         data.NewData(db),
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
