package app

import (
	"encoding/gob"
	"net/http"
	"os"
	"time"

	"github.com/lyticaa/lyticaa-app/internal/app/helpers"
	"github.com/lyticaa/lyticaa-app/internal/app/types"
	"github.com/lyticaa/lyticaa-app/internal/models"

	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/memcachier/mc"
	"github.com/newrelic/go-agent"
	"github.com/rs/zerolog/log"
	"gopkg.in/boj/redistore.v1"
)

type (
	App        types.App
	HTTP       types.HTTP
	Data       types.Data
	Monitoring types.Monitoring
)

func NewApp() *App {
	gob.Register(map[string]interface{}{})
	gob.Register(types.Flash{})
	gob.Register(types.Config{})
	gob.Register(models.UserModel{})

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

	sessionStore, err := redistore.NewRediStore(
		10,
		"tcp",
		os.Getenv("REDIS_URL"),
		os.Getenv("REDIS_PASSWORD"),
		[]byte(os.Getenv("SESSION_KEY")),
	)
	if err != nil {
		panic(err)
	}

	cache := mc.NewMC(
		os.Getenv("MEMCACHIER_SERVERS"),
		os.Getenv("MEMCACHIER_USERNAME"),
		os.Getenv("MEMCACHIER_PASSWORD"),
	)

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	var secure bool
	if helpers.Development() {
		secure = false
	} else {
		secure = true
	}

	return &App{
		Secure: secure,
		HTTP: types.HTTP{
			Router: mux.NewRouter(),
			Client: &http.Client{Timeout: 5 * time.Second},
		},
		Data: types.Data{
			SessionStore: sessionStore,
			Cache:        cache,
			Db:           db,
		},
		Monitoring: types.Monitoring{
			Logger:   log.With().Str("module", os.Getenv("APP_NAME")).Logger(),
			NewRelic: newRelic,
		},
	}
}
