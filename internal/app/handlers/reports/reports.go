package reports

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Reports struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewReports(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Reports {
	return &Reports{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
