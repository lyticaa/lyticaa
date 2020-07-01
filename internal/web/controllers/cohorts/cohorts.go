package cohorts

import (
	"gitlab.com/getlytica/lytica-app/internal/web/lib/data"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

const (
	highMargin     = "high_margin"
	lowMargin      = "low_margin"
	negativeMargin = "negative_margin"
)

type Cohorts struct {
	data         *data.Data
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewCohorts(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Cohorts {
	return &Cohorts{
		data:         data.NewData(db),
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
