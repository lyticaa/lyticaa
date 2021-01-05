package types

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/memcachier/mc"
	newrelic "github.com/newrelic/go-agent"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type HTTP struct {
	Router *mux.Router
	Server *http.Server
	Client *http.Client
}

type Database struct {
	Redis    *redistore.RediStore
	Memcache *mc.Client
	PG       *sqlx.DB
}

type Monitoring struct {
	Logger   zerolog.Logger
	NewRelic newrelic.Application
}

type App struct {
	Secure     bool
	HTTP       HTTP
	Database   Database
	Monitoring Monitoring
}
