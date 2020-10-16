package setup

import (
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/payments"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Setup struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
	stripe       payments.StripeGateway
}

func NewSetup(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger, stripe payments.StripeGateway) *Setup {
	return &Setup{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
		stripe:       stripe,
	}
}
