package helpers

import (
	"os"
)

const (
	development = "development"
	production  = "production"
)

func Development() bool {
	return os.Getenv("ENV") == development
}

func Production() bool {
	return os.Getenv("ENV") == production
}
