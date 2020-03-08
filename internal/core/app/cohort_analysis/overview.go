package cohort_analysis

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/helpers"
)

func (c *CohortAnalysis) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(c.sessionStore, c.logger, w, r)

	t := []string{"partials/nav/_main", "cohort_analysis/overview", "partials/_filters"}
	helpers.RenderTemplate(w, t, session.Values)
}
