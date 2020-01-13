package app

import (
	"context"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/handlers"
	"github.com/urfave/negroni"
)

var baseTmpl = "app"

func (a *App) Start() {
	a.Logger.Info().Msgf("starting on %v....", ":"+os.Getenv("PORT"))
	a.Router.Use(a.forceSsl)

	a.Handlers()
	a.RestHandlers()
	a.ErrorHandlers()

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

func (a *App) Handlers() {
	a.Router.HandleFunc("/", a.Home)
	a.Router.HandleFunc("/auth/login", a.Login)
	a.Router.HandleFunc("/auth/logout", a.Logout)
	a.Router.HandleFunc("/auth/callback", a.Callback)
	a.Router.Handle("/user", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.User)),
	))
	a.Router.Handle("/user/change_password", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.UserChangePassword)),
	))
	a.Router.Handle("/account/subscribe", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.AccountSubscribe)),
	))

	a.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/dist"))))
}

func (a *App) RestHandlers() {
	a.Router.HandleFunc("/api/health_check", a.HealthCheck)
}

func (a *App) ErrorHandlers() {
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

func (a *App) RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	cwd, _ := os.Getwd()
	t, err := template.ParseFiles(filepath.Join(cwd, "./web/templates/"+tmpl+".html"), filepath.Join(cwd, "./web/dist/"+baseTmpl+".html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, baseTmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
