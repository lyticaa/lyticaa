package helpers

import (
	"github.com/jmoiron/sqlx"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/models"

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

func ReloadSessionUser(session *sessions.Session, w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	SetSessionUser(*models.LoadUser(GetSessionUser(session).UserID, db), session, w, r)
}

func GetSessionUser(session *sessions.Session) models.User {
	user := session.Values["User"].(models.User)
	if user.Impersonate != nil && user.Admin {
		return *user.Impersonate
	}

	return user
}

func SetSessionUser(user models.User, session *sessions.Session, w http.ResponseWriter, r *http.Request) {
	var pUser models.User

	pUser = session.Values["User"].(models.User)
	if user.UserID != pUser.UserID {
		pUser.Impersonate = &user
	} else {
		pUser = user
	}

	session.Values["User"] = pUser
	_ = session.Save(r, w)
}

func IsSubscribed(session *sessions.Session) bool {
	subscribed := false

	user := GetSessionUser(session)
	if user.StripeSubscriptionID.Valid {
		subscribed = true
	}

	return subscribed
}
