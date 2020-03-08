package account

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Account struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewAccount(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Account {
	return &Account{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
