package helpers

import (
	"os"
)

const (
	development = "development"
)

func Development() bool {
	return os.Getenv("ENV") == development
}
