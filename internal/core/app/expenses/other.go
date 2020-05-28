package expenses

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (e *Expenses) Other(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)

	t := []string{
		"partials/nav/_main",
		"expenses/other",
		"partials/_filters",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (e *Expenses) OtherByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)
	_ = session.Values["User"].(models.User)

	table := []types.ExpensesTable{}
	byDate := types.Expenses{Data: table}

	js, err := json.Marshal(byDate)
	if err != nil {
		e.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
