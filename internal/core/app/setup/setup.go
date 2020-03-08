package setup

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Setup struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewSetup(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Setup {
	return &Setup{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
