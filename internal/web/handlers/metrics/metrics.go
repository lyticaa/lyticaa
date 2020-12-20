package metrics

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Metrics struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewMetrics(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Metrics {
	return &Metrics{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
