package core

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"time"

	"gitlab.com/getlytica/dashboard/internal/core/types"

	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/newrelic/go-agent"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/boj/redistore.v1"
)

type Core struct {
	Logger       zerolog.Logger
	NewRelic     newrelic.Application
	Srv          *http.Server
	Router       *mux.Router
	Client       *http.Client
	SessionStore *redistore.RediStore
	Db           *gorm.DB
}

func NewCore() *Core {
	gob.Register(map[string]interface{}{})

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

	sessionStore, err := redistore.NewRediStore(10, "tcp", os.Getenv("REDIS_URL"),
		os.Getenv("REDIS_PASSWORD"), []byte(os.Getenv("SESSION_KEY")))
	if err != nil {
		panic(err)
	}

	dbStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open("postgres", dbStr)
	if err != nil {
		panic(err)
	}

	return &Core{
		Logger:       log.With().Str("module", types.AppName).Logger(),
		NewRelic:     nr,
		Router:       mux.NewRouter(),
		Client:       &http.Client{Timeout: 5 * time.Second},
		SessionStore: sessionStore,
		Db:           db,
	}
}
