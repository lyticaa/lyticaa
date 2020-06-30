package cohorts

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (c *Cohorts) NegativeMargin(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(c.sessionStore, c.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"cohorts/negative_margin",
		"partials/_filters",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (c *Cohorts) NegativeMarginByDate(w http.ResponseWriter, r *http.Request) {
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
