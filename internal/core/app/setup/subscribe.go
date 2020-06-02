package setup

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/core/payments"
	"gitlab.com/getlytica/lytica-app/internal/models"
	"net/http"

	"github.com/gorilla/sessions"
)

func (s *Setup) stripeSessions(w http.ResponseWriter, session *sessions.Session) {
	user := session.Values["User"].(models.User)

	monthly, err := payments.CheckoutSession(user.UserId, user.Email, "monthly")
	if err != nil {
		s.logger.Error().Err(err).Msg("unable to generate a new stripe session")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	annual, err := payments.CheckoutSession(user.UserId, user.Email, "annual")
	if err != nil {
		s.logger.Error().Err(err).Msg("unable to generate a new stripe session")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	session.Values["stripeMonthlyId"] = monthly.ID
	session.Values["stripeAnnualId"] = annual.ID
}

func (s *Setup) Subscribe(w http.ResponseWriter, r *http.Request) {
	if helpers.IsSubscribed(s.sessionStore, s.logger, w, r) {
		http.Redirect(w, r, "/setup/subscribe/success", http.StatusSeeOther)
	}

	session := helpers.GetSession(s.sessionStore, s.logger, w, r)
	s.stripeSessions(w, session)

	session.Values["showPlans"] = true

	t := []string{
		"partials/_nav",
		"partials/nav/_setup",
		"partials/nav/account/_setup",
		"setup/subscribe",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (s *Setup) SubscribeSuccess(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(s.sessionStore, s.logger, w, r)
	session.Values["Flash"] = types.Flash{
		Success: types.FlashMessages["setup"]["subscribe"]["success"],
	}
	session.Values["allowNext"] = true
	session.Values["isSubscribed"] = true

	err := session.Save(r, w)
	if err != nil {
		s.logger.Error().Err(err).Msg("unable to save session")
	}

	t := []string{
		"partials/_nav",
		"partials/nav/_setup",
		"partials/nav/account/_setup",
		"partials/_flash",
		"setup/subscribe",
	}

	helpers.RenderTemplate(w, t, session.Values)
	helpers.ClearFlash(session, r, w)
}

func (s *Setup) SubscribeCancel(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(s.sessionStore, s.logger, w, r)
	s.stripeSessions(w, session)

	session.Values["Flash"] = types.Flash{
		Error: types.FlashMessages["setup"]["subscribe"]["error"],
	}

	t := []string{
		"partials/_nav",
		"partials/nav/_setup",
		"partials/nav/account/_setup",
		"partials/_flash",
		"setup/subscribe",
	}
	helpers.RenderTemplate(w, t, session.Values)
	helpers.ClearFlash(session, r, w)
}
