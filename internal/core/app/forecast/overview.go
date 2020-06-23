package forecast

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (f *Forecast) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(f.sessionStore, f.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"forecast/overview",
		"partials/_filters",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (f *Forecast) ForecastByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(f.sessionStore, f.logger, w, r)
	_ = session.Values["User"].(models.User)

	chart := types.Chart{}
	byDate := types.Forecast{Chart: chart}

	js, err := json.Marshal(byDate)
	if err != nil {
		f.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
