package cohorts

import (
	"encoding/json"
	"net/http"

	"gitlab.com/lyticaa/lyticaa-app/internal/web/helpers"
	"gitlab.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (c *Cohorts) NegativeMargin(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(c.sessionStore, c.logger, w, r)
	helpers.RenderTemplate(w, helpers.TemplateList(helpers.CohortsNegativeMargin), session.Values)
}

func (c *Cohorts) NegativeMarginByDate(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(c.sessionStore, c.logger, w, r))

	params := mux.Vars(r)
	dateRange := params["dateRange"]

	ok, _ := helpers.ValidateInput(helpers.ValidateDateRange{DateRange: dateRange}, &c.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var byDate types.Cohort
	byDate.Draw = helpers.DtDraw(r)

	c.data.Cohorts(user.UserId, dateRange, negativeMargin, &byDate, helpers.BuildFilter(r))
	js, err := json.Marshal(byDate)
	if err != nil {
		c.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
