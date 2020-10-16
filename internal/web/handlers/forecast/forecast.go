package forecast

import (
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/data"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Forecast struct {
	data         *data.Data
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewForecast(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Forecast {
	return &Forecast{
		data:         data.NewData(db),
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
