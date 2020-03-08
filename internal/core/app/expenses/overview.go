package expenses

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/helpers"
)

func (e *Expenses) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)

	t := []string{"partials/nav/_main", "expenses/overview", "partials/_filters"}
	helpers.RenderTemplate(w, t, session.Values)
}
