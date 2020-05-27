package cohorts

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
)

func (c *Cohorts) NegativeMargin(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(c.sessionStore, c.logger, w, r)

	t := []string{"partials/nav/_main", "cohorts/negative_margin", "partials/_filters"}
	helpers.RenderTemplate(w, t, session.Values)
}
