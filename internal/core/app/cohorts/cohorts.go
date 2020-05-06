package cohorts

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Cohorts struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewCohorts(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Cohorts {
	return &Cohorts{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
