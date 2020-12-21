package helpers

import (
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/models"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

func GetSession(store *redistore.RediStore, logger zerolog.Logger, w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := store.Get(r, "auth-session")
	if err != nil {
		logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resetFlash(session)
	setCSRF(session, r)

	return session
}

func GetSessionUser(session *sessions.Session) *models.UserModel {
	user := session.Values["User"].(models.UserModel)
	if user.Impersonate != nil && user.Admin {
		return user.Impersonate
	}

	return &user
}

func SetSessionUser(user *models.UserModel, session *sessions.Session, w http.ResponseWriter, r *http.Request) {
	var pUser models.UserModel

	pUser = session.Values["User"].(models.UserModel)
	if user.UserID != pUser.UserID {
		pUser.Impersonate = user
	} else {
		pUser = *user
	}

	session.Values["User"] = pUser
	_ = session.Save(r, w)
}

func resetFlash(session *sessions.Session) {
	session.Values["Flash"] = nil
}

func setCSRF(session *sessions.Session, r *http.Request) {
	session.Values[csrf.TemplateTag] = csrf.TemplateField(r)
}
