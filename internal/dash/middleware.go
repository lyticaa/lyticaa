package dash

import "net/http"

func (d *Dash) IsAuthenticated(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session, err := d.SessionStore.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := session.Values["profile"]; !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		next(w, r)
	}
}
