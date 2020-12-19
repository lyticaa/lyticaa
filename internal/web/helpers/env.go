package helpers

import (
	"os"
)

const (
	production = "production"
)

func Production() bool {
	if os.Getenv("ENV") == production {
		return true
	}

	return false
}
