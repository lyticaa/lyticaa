package app

import (
	"os"

	"github.com/getsentry/sentry-go"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/newrelic/go-agent"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type App struct {
	Logger   zerolog.Logger
	NewRelic newrelic.Application
	Db       *sqlx.DB
}

func NewApp() *App {
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
	nr, _ := newrelic.NewApplication(config)

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	return &App{
		Logger:   log.With().Str("module", os.Getenv("APP_NAME")).Logger(),
		NewRelic: nr,
		Db:       db,
	}
}
