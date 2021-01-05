package ingest

import (
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/newrelic/go-agent"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Ingest struct {
	Logger   zerolog.Logger
	NewRelic newrelic.Application
	Db       *sqlx.DB
}

func NewIngest(nr newrelic.Application, db *sqlx.DB) *Ingest {
	return &Ingest{
		Logger:   log.With().Str("module", os.Getenv("APP_NAME")).Logger(),
		NewRelic: nr,
		Db:       db,
	}
}
