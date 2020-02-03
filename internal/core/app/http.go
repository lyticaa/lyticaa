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

var (
	cwd, _   = os.Getwd()
	baseTmpl = "app"
	baseFiles = []string{
		filepath.Join(cwd, "./web/dist/"+baseTmpl+".html"),
		filepath.Join(cwd, "./web/templates/partials/_nav.gohtml"),
		filepath.Join(cwd, "./web/templates/partials/_footer.gohtml"),
	}
)

func (a *App) Start() {
	a.Logger.Info().Msgf("starting on %v....", ":"+os.Getenv("PORT"))
	a.Router.Use(a.forceSsl)

	a.handlers()
	a.restHandlers()
	a.errorHandlers()

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

func (a *App) handlers() {
	a.Router.HandleFunc("/auth/login", a.login)
	a.Router.HandleFunc("/auth/logout", a.logout)
	a.Router.HandleFunc("/auth/callback", a.callback)

	a.Router.Handle("/", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.home)),
	))
	a.Router.Handle("/metrics/total_sales", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.totalSales)),
	))
	a.Router.Handle("/metrics/units_sold", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.unitsSold)),
	))
	a.Router.Handle("/metrics/amazon_costs", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.amazonCosts)),
	))
	a.Router.Handle("/metrics/advertising", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.advertising)),
	))
	a.Router.Handle("/metrics/refunds", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.refunds)),
	))
	a.Router.Handle("/metrics/shipping_credits", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.shippingCredits)),
	))
	a.Router.Handle("/metrics/promotional_rebates", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.promotionalRebates)),
	))
	a.Router.Handle("/metrics/total_costs", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.totalCosts)),
	))
	a.Router.Handle("/metrics/net_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.netMargin)),
	))

	a.Router.Handle("/cohort_analysis", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.cohortAnalysis)),
	))

	a.Router.Handle("/forecast", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.forecast)),
	))

	a.Router.Handle("/expenses", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.expenses)),
	))

	a.Router.Handle("/profit_loss", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.profitLoss)),
	))

	a.Router.Handle("/user/notifications", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.notifications)),
	))
	a.Router.Handle("/user/invitations", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.invitations)),
	))
	a.Router.Handle("/user/subscription", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.subscription)),
	))
	a.Router.Handle("/user/change_password", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.changePassword)),
	))
	a.Router.Handle("/user/subscribe", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.subscribe)),
	))

	a.Router.Handle("/onboard/details", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.details)),
	))
	a.Router.Handle("/onboard/team", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.team)),
	))
	a.Router.Handle("/onboard/subscribe", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.subscribe)),
	))
	a.Router.Handle("/onboard/import_data", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.importData)),
	))
	a.Router.Handle("/onboard/complete", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(a.complete)),
	))

	a.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/dist"))))
}

func (a *App) restHandlers() {
	a.Router.HandleFunc("/api/health_check", a.healthCheck)
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

func (a *App) templateList(fileList []string) []string {
	var container []string
	container = append(container, baseFiles...)

	for _, file := range fileList {
		container = append(container, filepath.Join(cwd, "./web/templates/"+file+".gohtml"))
	}

	return container
}

func (a *App) renderTemplate(w http.ResponseWriter, templates []string, data interface{}) {
	files := a.templateList(templates)
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.ExecuteTemplate(w, baseTmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
