package main

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

func main() {
	dbStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	debug := log.With().Str("module", os.Getenv("APP_NAME")).Logger()

	m, err := migrate.New("file://./db/migrations", dbStr)
	if err != nil {
		debug.Error().Err(err).Msg("Error")
		return
	}
	defer m.Close()

	err = m.Up()
	if err != nil && err != m.ErrNoChange {
		debug.Error().Err(err).Msg("Error")
		return
	}
}
