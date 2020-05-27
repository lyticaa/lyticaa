package dashboard

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Dashboard struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewDashboard(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Dashboard {
	return &Dashboard{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
