package profit_loss

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/helpers"
)

func (p *ProfitLoss) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(p.sessionStore, p.logger, w, r)

	t := []string{"partials/nav/_main", "profit_loss/overview", "partials/_filters"}
	helpers.RenderTemplate(w, t, session.Values)
}
