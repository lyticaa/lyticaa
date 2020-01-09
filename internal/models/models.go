package models

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func logger() *zerolog.Logger {
	log := log.With().Str("module", os.Getenv("APP_NAME")).Logger()
	return &log
}
