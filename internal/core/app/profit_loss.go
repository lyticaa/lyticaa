package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/profit_loss"

	"github.com/urfave/negroni"
)

func (a *App) profitLossHandlers() {
	p := profit_loss.NewProfitLoss(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/profit_loss", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(p.Overview)),
	))
}
