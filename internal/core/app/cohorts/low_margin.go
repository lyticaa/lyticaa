package cohorts

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/helpers"
)

func (c *Cohorts) LowMargin(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(c.sessionStore, c.logger, w, r)

	t := []string{"partials/nav/_main", "cohorts/low_margin", "partials/_filters"}
	helpers.RenderTemplate(w, t, session.Values)
}
