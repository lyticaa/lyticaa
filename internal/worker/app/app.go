package app

import (
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/newrelic/go-agent"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type App struct {
	Logger   zerolog.Logger
	NewRelic newrelic.Application
	Db       *gorm.DB
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
		os.Getenv("NEWRELIC_LICENSE_KEY"),
	)
	nr, _ := newrelic.NewApplication(config)
	//
	//dbStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v",
	//	os.Getenv("DB_HOST"),
	//	os.Getenv("DB_PORT"),
	//	os.Getenv("DB_USERNAME"),
	//	os.Getenv("DB_NAME"),
	//	os.Getenv("DB_PASSWORD"),
	//	os.Getenv("DB_SSLMODE"),
	//)
	//
	//db, err := gorm.Open("postgres", dbStr)
	//if err != nil {
	//	panic(err)
	//}

	return &App{
		Logger:   log.With().Str("module", os.Getenv("APP_NAME")).Logger(),
		NewRelic: nr,
		//Db:       db,
	}
}
