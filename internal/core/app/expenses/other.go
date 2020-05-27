package expenses

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
)

func (e *Expenses) Other(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)

	t := []string{"partials/nav/_main", "expenses/other", "partials/_filters"}
	helpers.RenderTemplate(w, t, session.Values)
}
