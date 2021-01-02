package forecast

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Forecast struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewForecast(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Forecast {
	return &Forecast{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
