package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Auth struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewAuth(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Auth {
	return &Auth{
		db:           db,
		sessionStore: sessionStore,
		logger:       log,
	}
}
