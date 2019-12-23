package dash

import (
	"net/http"
)

func (d *Dash) Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/home.html")
}
