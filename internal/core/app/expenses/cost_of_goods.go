package expenses

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
)

func (e *Expenses) CostOfGoods(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)

	t := []string{"partials/nav/_main", "expenses/cost_of_goods", "partials/_filters"}
	helpers.RenderTemplate(w, t, session.Values)
}
