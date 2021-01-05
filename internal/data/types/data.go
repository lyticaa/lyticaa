package types

import (
	"github.com/jmoiron/sqlx"
	newrelic "github.com/newrelic/go-agent"
	"github.com/rs/zerolog"
)

type Database struct {
	PG *sqlx.DB
}

type Monitoring struct {
	Logger   zerolog.Logger
	NewRelic newrelic.Application
}

type Data struct {
	Database   Database
	Monitoring Monitoring
}
