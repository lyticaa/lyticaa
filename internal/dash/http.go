package dash

import (
	"context"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/negroni"
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
	d.Router.HandleFunc("/auth/login", d.Login)
	d.Router.HandleFunc("/auth/logout", d.Logout)
	d.Router.HandleFunc("/auth/callback", d.Callback)
	d.Router.Handle("/user", negroni.New(
		negroni.HandlerFunc(d.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(d.User)),
	))

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

func (d *Dash) RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(filepath.Join(cwd, "./web/templates/"+tmpl+".html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
