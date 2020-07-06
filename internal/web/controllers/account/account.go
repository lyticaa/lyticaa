package account

import (
	"gitlab.com/getlytica/lytica-app/internal/web/lib/payments"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Account struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
	stripe       payments.StripeGateway
}

func NewAccount(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger, stripe payments.StripeGateway) *Account {
	return &Account{
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
		stripe:       stripe,
	}
}
