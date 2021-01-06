package webhooks

import (
	"io/ioutil"
	"net/http"

	"github.com/lyticaa/lyticaa/internal/app/pkg/accounts/payments"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Webhooks struct {
	db     *sqlx.DB
	logger zerolog.Logger
	stripe payments.StripeGateway
}

func NewWebhooks(db *sqlx.DB, log zerolog.Logger, stripe payments.StripeGateway) *Webhooks {
	return &Webhooks{
		db:     db,
		logger: log,
		stripe: stripe,
	}
}

func (wh *Webhooks) readBody(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		return nil, err
	}

	return body, nil
}
