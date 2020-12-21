package home

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Home struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewHome(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Home {
	return &Home{
		db:           db,
		sessionStore: sessionStore,
		logger:       log,
	}
}
