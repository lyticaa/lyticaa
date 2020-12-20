package app

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/csrf"
)

func (a *App) Start() {
	a.Monitoring.Logger.Info().Msgf("starting on %v....", ":"+os.Getenv("PORT"))
	a.HTTP.Router.Use(a.ForceSSL)

	a.initializeHandlers()

	a.HTTP.Server = &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      csrf.Protect([]byte(os.Getenv("CSRF_TOKEN")))(a.HTTP.Router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		if err := a.HTTP.Server.ListenAndServe(); err != nil {
			a.Monitoring.Logger.Info().Err(err)
		}
	}()
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.HTTP.Server.Shutdown(ctx); err != nil {
		a.Monitoring.Logger.Fatal().Err(err).Msg("server shutdown")
	}

	a.Monitoring.Logger.Info().Msg("server exiting....")
}

func (a *App) initializeHandlers() {
	a.accountHandlers()
	a.adminHandlers()
	a.apiHandlers()
	a.authHandlers()
	a.cohortsHandlers()
	a.dashboardHandlers()
	a.expensesHandlers()
	a.forecastHandlers()
	a.homeHandlers()
	a.metricsHandlers()
	a.profitLossHandlers()
	a.reportsHandlers()
	a.webhookHandlers()

	a.HTTP.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/dist"))))
}
