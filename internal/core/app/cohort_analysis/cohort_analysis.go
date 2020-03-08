package cohort_analysis

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type CohortAnalysis struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewCohortAnalysis(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *CohortAnalysis {
	return &CohortAnalysis{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
