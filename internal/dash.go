package dash

import (
	"net/http"
	"os"
	"time"

	"gitlab.com/sellernomics/dashboard/internal/types"

	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Dash struct {
	Logger   zerolog.Logger
	NewRelic newrelic.Application
	Srv      *http.Server
	Router   *mux.Router
	Client   *http.Client
}

func NewDash() *Dash {
	sentryOpts := sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}
	err := sentry.Init(sentryOpts)
	if err != nil {
		panic(err)
	}

	config := newrelic.NewConfig(
		os.Getenv("APP_NAME"),
		os.Getenv("NEWRELIC_LICENSE_KEY"),
	)
	nr, _ := newrelic.NewApplication(config)

	return &Dash{
		Logger:   log.With().Str("module", types.AppName).Logger(),
		NewRelic: nr,
		Router:   mux.NewRouter(),
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
