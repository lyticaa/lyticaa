package metrics

import (
	"encoding/json"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (m *Metrics) ShippingCredits(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)
	helpers.RenderTemplate(w, helpers.TemplateList(helpers.MetricsShippingCredits), session.Values)
}

func (m *Metrics) ShippingCreditsByDate(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(m.sessionStore, m.logger, w, r))

	params := mux.Vars(r)
	dateRange := params["dateRange"]

	ok, _ := helpers.ValidateInput(helpers.ValidateDateRange{DateRange: dateRange}, &m.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var byDate types.ShippingCredits
	byDate.Draw = helpers.DtDraw(r)

	m.data.MetricsShippingCredits(user.UserId, dateRange, &byDate, helpers.BuildFilter(r))
	js, err := json.Marshal(byDate)
	if err != nil {
		m.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
