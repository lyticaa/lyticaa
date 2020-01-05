package core

import (
	"context"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/urfave/negroni"
)

func (c *Core) Start() {
	c.Logger.Info().Msgf("starting on %v....", ":"+os.Getenv("PORT"))
	c.Router.Use(c.forceSsl)

	c.Handlers()
	c.RestHandlers()

	c.Srv = &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: c.Router,
	}

	go func() {
		if err := c.Srv.ListenAndServe(); err != nil {
			c.Logger.Info().Err(err)
		}
	}()
}

func (c *Core) Handlers() {
	c.Router.HandleFunc("/", c.Home)
	c.Router.HandleFunc("/auth/login", c.Login)
	c.Router.HandleFunc("/auth/logout", c.Logout)
	c.Router.HandleFunc("/auth/callback", c.Callback)
	c.Router.Handle("/user", negroni.New(
		negroni.HandlerFunc(c.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(c.User)),
	))
	c.Router.Handle("/account/subscribe", negroni.New(
		negroni.HandlerFunc(c.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(c.AccountSubscribe)),
	))

	c.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
}

func (c *Core) RestHandlers() {
	c.Router.HandleFunc("/api/v1/health_check", c.HealthCheck)
}

func (c *Core) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.Srv.Shutdown(ctx); err != nil {
		c.Logger.Fatal().Err(err).Msg("server shutdown")
	}

	c.Logger.Info().Msg("server exiting....")
}

func (c *Core) RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
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
