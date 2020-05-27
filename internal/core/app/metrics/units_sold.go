package metrics

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
)

func (m *Metrics) UnitsSold(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)

	t := []string{"partials/nav/_main", "metrics/units_sold", "partials/_filters"}
	helpers.RenderTemplate(w, t, session.Values)
}
