package setup

import (
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/sessions"
)

func (s *Setup) Subscribe(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(s.sessionStore, s.logger, w, r)
	if helpers.IsSubscribed(session) {
		http.Redirect(w, r, "/setup/subscribe/success", http.StatusSeeOther)
	}

	if err := s.stripeSessions(w, session); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	helpers.RenderTemplate(w, helpers.TemplateList(helpers.SetupSubscribe), session.Values)
}

func (s *Setup) SubscribeSuccess(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(s.sessionStore, s.logger, w, r)
	session.Values["Flash"] = types.Flash{
		Success: types.FlashMessages["setup"]["subscribe"]["success"],
	}

	helpers.ReloadSessionUser(session, w, r, s.db)

	err := session.Save(r, w)
	if err != nil {
		s.logger.Error().Err(err).Msg("unable to save session")
	}

	helpers.RenderTemplate(w, helpers.TemplateList(helpers.SetupSubscribe), session.Values)
	helpers.ClearFlash(session, r, w)
}

func (s *Setup) SubscribeCancel(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(s.sessionStore, s.logger, w, r)
	if err := s.stripeSessions(w, session); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	session.Values["Flash"] = types.Flash{
		Error: types.FlashMessages["setup"]["subscribe"]["error"],
	}

	helpers.RenderTemplate(w, helpers.TemplateList(helpers.SetupSubscribe), session.Values)
	helpers.ClearFlash(session, r, w)
}

func (s *Setup) stripeSessions(w http.ResponseWriter, session *sessions.Session) error {
	user := helpers.GetSessionUser(session)
	stripeSessions, err := s.stripe.CheckoutSessions(user.UserID, user.Email)
	if err != nil {
		return err
	}

	session.Values["stripeMonthlyID"] = (*stripeSessions)[0]
	session.Values["stripeAnnualID"] = (*stripeSessions)[1]

	return nil
}
