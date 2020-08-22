package webhooks

import (
	"io/ioutil"
	"net/http"

	"gitlab.com/lyticaa/lyticaa-app/internal/web/lib/payments"

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
		wh.logger.Error().Err(err).Msg("failed to read the webhook body")
		w.WriteHeader(http.StatusServiceUnavailable)
		return nil, err
	}

	return body, nil
}
