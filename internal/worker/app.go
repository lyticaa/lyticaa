package app

import (
	"context"
	"os"
	"os/signal"

	"github.com/getsentry/sentry-go"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/memcachier/mc"
	"github.com/newrelic/go-agent"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Signalling struct {
	ctx    context.Context
	cancel context.CancelFunc
	stop   chan struct{}
	quit   chan os.Signal
}

type App struct {
	Logger   zerolog.Logger
	NewRelic newrelic.Application
	Db       *sqlx.DB
	Cache    *mc.Client
	Signalling
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

	cache := mc.NewMC(os.Getenv("MEMCACHED_SERVERS"),
		os.Getenv("MEMCACHED_USERNAME"),
		os.Getenv("MEMCACHED_PASSWORD"))

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	return &App{
		Logger:   log.With().Str("app", os.Getenv("APP_NAME")).Str("component", "worker").Logger(),
		NewRelic: nr,
		Db:       db,
		Cache:    cache,
		Signalling: Signalling{
			ctx:    ctx,
			cancel: cancel,
			stop:   make(chan struct{}),
			quit:   quit,
		},
	}
}
