package home

import (
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Home struct {
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewHome(sessionStore *redistore.RediStore, log zerolog.Logger) *Home {
	return &Home{
		sessionStore: sessionStore,
		logger:       log,
	}
}
