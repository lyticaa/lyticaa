package forecast

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
)

func (f *Forecast) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(f.sessionStore, f.logger, w, r)

	t := []string{"partials/nav/_main", "forecast/overview", "partials/_filters"}
	helpers.RenderTemplate(w, t, session.Values)
}
