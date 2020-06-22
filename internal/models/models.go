package models

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Filter struct {
	DateRange string
	Start     int64
	Length    int64
	Sort      int64
	Dir       string
	StartDate time.Time
	EndDate   time.Time
}

func NewFilter() *Filter {
	return &Filter{}
}

func logger() *zerolog.Logger {
	log := log.With().Str("module", os.Getenv("APP_NAME")).Logger()
	return &log
}

func sortColumn(columnMap map[int64]string, columnIdx int64) string {
	if columnIdx > int64(len(columnMap)) {
		return columnMap[0]
	}

	return columnMap[columnIdx]
}
