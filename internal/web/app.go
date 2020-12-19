package app

import (
	"encoding/gob"
	"net/http"
	"os"
	"time"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/memcachier/mc"
	"github.com/newrelic/go-agent"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/boj/redistore.v1"
)

type HTTP struct {
	Router *mux.Router
	Server *http.Server
	Client *http.Client
}

type Data struct {
	SessionStore *redistore.RediStore
	Cache        *mc.Client
	Db           *sqlx.DB
}

type Monitoring struct {
	Logger   zerolog.Logger
	NewRelic newrelic.Application
}

type App struct {
	HTTP       HTTP
	Data       Data
	Monitoring Monitoring
}

func NewApp() *App {
	gob.Register(map[string]interface{}{})
	gob.Register(types.Flash{})
	gob.Register(types.Config{})
	gob.Register(models.User{})

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
		os.Getenv("MEMCACHED_SERVERS"),
		os.Getenv("MEMCACHED_USERNAME"),
		os.Getenv("MEMCACHED_PASSWORD"),
	)

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	return &App{
		HTTP: HTTP{
			Router: mux.NewRouter(),
			Client: &http.Client{Timeout: 5 * time.Second},
		},
		Data: Data{
			SessionStore: sessionStore,
			Cache:        cache,
			Db:           db,
		},
		Monitoring: Monitoring{
			Logger:   log.With().Str("module", os.Getenv("APP_NAME")).Logger(),
			NewRelic: newRelic,
		},
	}
}
