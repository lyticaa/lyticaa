package admin

import (
	"encoding/json"
	"net/http"
	"os"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (a *Admin) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := session.Values["User"].(models.User)

	if !user.Admin {
		http.Redirect(w, r, os.Getenv("BASE_URL"), 302)
	}

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_account",
		"partials/nav/account/_main",
		"partials/filters/_filters",
		"admin/overview",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (a *Admin) UsersByDate(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(a.paintUsers(r))
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (a *Admin) Impersonate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := session.Values["User"].(models.User)

	params := mux.Vars(r)
	user.Impersonate = models.LoadUser(params["user"], a.db)
	session.Values["User"] = user
	_ = session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *Admin) LogOut(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := session.Values["User"].(models.User)

	session.Values["User"] = models.LoadUser(user.UserId, a.db)
	_ = session.Save(r, w)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (a *Admin) paintUsers(r *http.Request) types.Admin {
	filter := helpers.BuildFilter(r)
	users := models.LoadUsers(filter, a.db)
	var byDate types.Admin

	for _, user := range *users {
		t := types.AdminTable{
			RowId:   user.UserId,
			Email:   user.Email,
			Created: user.CreatedAt.Format("2006-01-02"),
		}

		byDate.Data = append(byDate.Data, t)
	}

	if len(byDate.Data) == 0 {
		byDate.Data = []types.AdminTable{}
	}

	byDate.Draw = helpers.DtDraw(r)
	byDate.RecordsTotal = models.TotalUsers(filter, a.db)
	byDate.RecordsFiltered = byDate.RecordsTotal

	return byDate
}
