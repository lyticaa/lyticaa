package metrics

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (m *Metrics) NetMargin(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_account",
		"partials/nav/account/_main",
		"partials/admin/_impersonate",
		"partials/filters/_filters",
		"partials/filters/_date",
		"partials/filters/_import",
		"metrics/net_margin",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (m *Metrics) NetMarginByDate(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(m.sessionStore, m.logger, w, r))

	params := mux.Vars(r)
	dateRange := params["dateRange"]

	ok, _ := helpers.ValidateInput(helpers.ValidateDateRange{DateRange: dateRange}, &m.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var byDate types.NetMargin
	byDate.Draw = helpers.DtDraw(r)

	m.data.MetricsNetMargin(user.UserId, dateRange, &byDate, helpers.BuildFilter(r))
	js, err := json.Marshal(byDate)
	if err != nil {
		m.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
