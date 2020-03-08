package expenses

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Expenses struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewExpenses(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Expenses {
	return &Expenses{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
