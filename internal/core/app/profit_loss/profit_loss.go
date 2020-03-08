package profit_loss

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type ProfitLoss struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewProfitLoss(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *ProfitLoss {
	return &ProfitLoss{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}
