package helpers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

func resetFlash(session *sessions.Session) *sessions.Session {
	session.Values["Flash"] = nil
	return session
}

func GetSession(store *redistore.RediStore, logger zerolog.Logger, w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := store.Get(r, "auth-session")
	if err != nil {
		logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return resetFlash(session)
}

func IsSubscribed(sessionStore *redistore.RediStore, logger zerolog.Logger, w http.ResponseWriter, r *http.Request) bool {
	ok := false

	session := GetSession(sessionStore, logger, w, r)
	if session.Values["isSubscribed"] == nil {
		return ok
	}

	subscribed := session.Values["isSubscribed"].(bool)
	if subscribed {
		ok = true
	}

	return ok
}
