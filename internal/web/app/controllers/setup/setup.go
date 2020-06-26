package setup

import (
	"gitlab.com/getlytica/lytica-app/internal/web/app/lib/payments"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Setup struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
	stripe       payments.Gateway
}

func NewSetup(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger, stripe payments.Gateway) *Setup {
	return &Setup{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
		stripe:       stripe,
	}
}
