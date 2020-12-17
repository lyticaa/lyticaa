package main

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

var ErrNoChange = "no change"

func main() {
	debug := log.With().Str("module", os.Getenv("APP_NAME")).Logger()

	m, err := migrate.New("file://./db/migrations", os.Getenv("DATABASE_URL"))
	if err != nil {
		debug.Error().Err(err).Msg("Error")
		return
	}

	defer m.Close()

	err = m.Up()
	if err != nil && err.Error() != ErrNoChange {
		debug.Error().Err(err).Msg("Error")
		return
	}
}
