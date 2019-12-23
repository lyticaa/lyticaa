package dash

import (
	"context"
	"net/http"
	"os"
	"time"
)

func (d *Dash) Start() {
	d.Logger.Info().Msgf("starting on %v....", ":"+os.Getenv("PORT"))
	d.Handlers()

	d.Srv = &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: d.Router,
	}

	go func() {
		if err := d.Srv.ListenAndServe(); err != nil {
			d.Logger.Info().Err(err)
		}
	}()
}

func (d *Dash) Handlers() {
	d.Router.HandleFunc("/", d.Home)
	d.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
}

func (d *Dash) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := d.Srv.Shutdown(ctx); err != nil {
		d.Logger.Fatal().Err(err).Msg("server shutdown")
	}

	d.Logger.Info().Msg("server exiting....")
}
