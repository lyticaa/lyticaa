package cohorts

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (c *Cohorts) LowMargin(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(c.sessionStore, c.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"partials/cohorts/_margin",
		"partials/filters/_filters",
		"partials/filters/_date",
		"partials/filters/_upload",
		"cohorts/low_margin",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (c *Cohorts) LowMarginByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(c.sessionStore, c.logger, w, r)
	user := session.Values["User"].(models.User)

	params := mux.Vars(r)
	dateRange := params["dateRange"]

	ok, _ := helpers.ValidateInput(helpers.ValidateDateRange{DateRange: dateRange}, &c.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var byDate types.Cohort
	byDate.Draw = helpers.DtDraw(r)

	c.data.Cohorts(user.UserId, dateRange, lowMargin, &byDate, helpers.BuildFilter(r))
	js, err := json.Marshal(byDate)
	if err != nil {
		c.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
