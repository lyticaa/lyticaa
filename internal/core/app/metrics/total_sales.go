package metrics

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/helpers"
)

func (m *Metrics) TotalSales(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)

	t := []string{"partials/nav/_main", "metrics/total_sales", "partials/_filters"}
	helpers.RenderTemplate(w, t, session.Values)
}
