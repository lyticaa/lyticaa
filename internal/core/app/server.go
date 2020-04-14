package app

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/urfave/negroni"
)

func (a *App) Start() {
	a.Logger.Info().Msgf("starting on %v....", ":"+os.Getenv("PORT"))
	a.Router.Use(a.forceSsl)

	a.initializeHandlers()

	a.Srv = &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      a.Router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		if err := a.Srv.ListenAndServe(); err != nil {
			a.Logger.Info().Err(err)
		}
	}()
}

func (a *App) initializeHandlers() {
	a.restHandlers()
	a.webhookHandlers()
	a.errorHandlers()
	a.accountHandlers()
	a.cohortAnalysisHandlers()
	a.dataHandlers()
	a.expensesHandlers()
	a.forecastHandlers()
	a.metricsHandlers()
	a.profitLossHandlers()
	a.setupHandlers()
	a.authHandlers()
}

func (a *App) authHandlers() {
	a.Router.HandleFunc("/auth/login", a.login)
	a.Router.HandleFunc("/auth/logout", a.logout)
	a.Router.HandleFunc("/auth/callback", a.callback)

	a.Router.Handle("/", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(a.home)),
	))

	a.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/dist"))))
}

func (a *App) restHandlers() {
	a.Router.Handle("/api/health_check", negroni.New(
		negroni.Wrap(http.HandlerFunc(a.healthCheck)),
	))
}

func (a *App) webhookHandlers() {
	a.Router.Handle("/webhooks/stripe", negroni.New(
		negroni.Wrap(http.HandlerFunc(a.stripeWebhooks)),
	))
}

func (a *App) errorHandlers() {
	a.Router.NotFoundHandler = handlers.LoggingHandler(
		os.Stdout,
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
		),
	)

	a.Router.MethodNotAllowedHandler = handlers.LoggingHandler(
		os.Stdout,
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusMethodNotAllowed)
			},
		),
	)
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.Srv.Shutdown(ctx); err != nil {
		a.Logger.Fatal().Err(err).Msg("server shutdown")
	}

	a.Logger.Info().Msg("server exiting....")
}
