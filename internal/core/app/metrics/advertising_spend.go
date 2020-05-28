package metrics

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (m *Metrics) AdvertisingSpend(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)

	t := []string{
		"partials/nav/_main",
		"metrics/advertising_spend",
		"partials/_filters",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (m *Metrics) AdvertisingSpendByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)
	_ = session.Values["User"].(models.User)

	table := []types.AdvertisingSpendTable{}
	byDate := types.AdvertisingSpend{Data: table}

	js, err := json.Marshal(byDate)
	if err != nil {
		m.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
