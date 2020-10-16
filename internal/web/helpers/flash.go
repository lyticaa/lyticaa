package helpers

import (
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/sessions"
)

func ClearFlash(session *sessions.Session, r *http.Request, w http.ResponseWriter) {
	session.Values["Flash"] = types.Flash{}
	_ = session.Save(r, w)
}

func SetFlash(flash, message string, session *sessions.Session, r *http.Request, w http.ResponseWriter) error {
	ClearFlash(session, r, w)

	switch flash {
	case "success":
		setFlashSuccess(message, session)
	case "error":
		setFlashError(message, session)
	case "warning":
		setFlashWarning(message, session)
	case "info":
		setFlashInfo(message, session)
	}

	err := session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func setFlashSuccess(message string, session *sessions.Session) {
	session.Values["Flash"] = types.Flash{
		Success: message,
	}
}

func setFlashError(message string, session *sessions.Session) {
	session.Values["Flash"] = types.Flash{
		Error: message,
	}
}

func setFlashWarning(message string, session *sessions.Session) {
	session.Values["Flash"] = types.Flash{
		Warning: message,
	}
}

func setFlashInfo(message string, session *sessions.Session) {
	session.Values["Flash"] = types.Flash{
		Info: message,
	}
}
