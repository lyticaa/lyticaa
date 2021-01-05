package data

import (
	"os"

	"github.com/lyticaa/lyticaa-app/internal/data/types"

	"github.com/getsentry/sentry-go"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/newrelic/go-agent"
	"github.com/rs/zerolog/log"
)

type (
	Data       types.Data
	Monitoring types.Monitoring
)

func NewData() *Data {
	sentryOpts := sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
	}
	err := sentry.Init(sentryOpts)
	if err != nil {
		panic(err)
	}

	config := newrelic.NewConfig(
		os.Getenv("APP_NAME"),
		os.Getenv("NEW_RELIC_LICENSE_KEY"),
	)
	newRelic, _ := newrelic.NewApplication(config)

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	return &Data{
		Database: types.Database{
			PG: db,
		},
		Monitoring: types.Monitoring{
			Logger:   log.With().Str("module", os.Getenv("APP_NAME")).Logger(),
			NewRelic: newRelic,
		},
	}
}
