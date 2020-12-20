package admin

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/admin"
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/users"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (a *Admin) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := helpers.GetSessionUser(helpers.GetSession(a.sessionStore, a.logger, w, r))

	if !user.Admin {
		http.Redirect(w, r, os.Getenv("BASE_URL"), 302)
	}

	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.AdminOverview), session.Values)
}

func (a *Admin) UsersByDate(w http.ResponseWriter, r *http.Request) {
	var adminList types.Admin
	admin.Users(r.Context(), &adminList, helpers.BuildFilter(r), a.db)

	js, err := json.Marshal(adminList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (a *Admin) Impersonate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := helpers.GetSessionUser(helpers.GetSession(a.sessionStore, a.logger, w, r))

	params := mux.Vars(r)

	impersonate := users.FetchUser(r.Context(), params["user"], a.db)
	user.Impersonate = &impersonate
	session.Values["User"] = user
	_ = session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *Admin) LogOut(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := helpers.GetSessionUser(helpers.GetSession(a.sessionStore, a.logger, w, r))

	session.Values["User"] = users.FetchUser(r.Context(), user.UserID, a.db)
	_ = session.Save(r, w)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}
