package helpers

import (
	"os"
)

const (
	production = "production"
)

func Production() bool {
	return os.Getenv("ENV") == production
}
