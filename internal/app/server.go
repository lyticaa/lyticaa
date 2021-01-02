package app

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/csrf"
)

func (a *App) Start(withCSRF bool) {
	a.Monitoring.Logger.Info().Msgf("starting on %v....", ":"+os.Getenv("PORT"))
	a.HTTP.Router.Use(a.Secured)

	var handler http.Handler
	if withCSRF {
		handler = csrf.Protect([]byte(os.Getenv("CSRF_TOKEN")), csrf.Secure(a.Secure))(a.HTTP.Router)
	} else {
		handler = a.HTTP.Router
	}

	a.HTTP.Server = &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      handler,
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
