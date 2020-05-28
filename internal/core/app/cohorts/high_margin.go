package cohorts

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (c *Cohorts) HighMargin(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(c.sessionStore, c.logger, w, r)

	t := []string{
		"partials/nav/_main",
		"cohorts/high_margin",
		"partials/_filters",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (c *Cohorts) HighMarginByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(c.sessionStore, c.logger, w, r)
	_ = session.Values["User"].(models.User)

	table := []types.CohortTable{}
	byDate := types.Cohort{Data: table}

	js, err := json.Marshal(byDate)
	if err != nil {
		c.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
