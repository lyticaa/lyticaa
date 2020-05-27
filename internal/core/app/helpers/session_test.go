package helpers

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"

	"gitlab.com/getlytica/lytica/internal/core/app/types"
	"gitlab.com/getlytica/lytica/internal/models"

	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

const (
	sessionName = "auth-session"
)

func TestGetSession(t *testing.T) {
	var log zerolog.Logger

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error()
	}

	s := GetSession(testSessionStore(t), log, nil, r)
	if s.Name() != sessionName {
		t.Error()
	}
}

func testSessionStore(t *testing.T) *redistore.RediStore {
	gob.Register(map[string]interface{}{})
	gob.Register(types.Flash{})
	gob.Register(types.Config{})
	gob.Register(models.User{})

	store, err := redistore.NewRediStore(
		10,
		"tcp",
		os.Getenv("REDIS_URL"),
		os.Getenv("REDIS_PASSWORD"),
		[]byte(os.Getenv("SESSION_KEY")))
	if err != nil {
		t.Error(err)
	}

	return store
}
