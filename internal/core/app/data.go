package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/data"

	"github.com/urfave/negroni"
)

func (a *App) dataHandlers() {
	a.Router.Handle("/upload", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(data.Upload)),
	))
}
