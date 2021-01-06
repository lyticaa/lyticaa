package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lyticaa/lyticaa/internal/app/helpers"
	"github.com/lyticaa/lyticaa/internal/app/pkg/accounts"
	"github.com/lyticaa/lyticaa/internal/app/types"

	"github.com/gorilla/mux"
)

const (
	successResult = "success"
	cancelResult  = "cancel"
)

func (a *Account) Subscription(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := helpers.GetSessionUser(session)

	subscription := accounts.Subscription(r.Context(), user.ID, a.db)
	if subscription.ID != 0 {
		session.Values["Subscription"] = subscription
	} else {
		monthly, annual, err := accounts.NewStripeSessions(user.UserID, user.Email, session.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		session.Values["StripeMonthlyID"] = *monthly
		session.Values["StripeAnnualID"] = *annual
	}

	if mux.Vars(r)["sessionID"] == session.ID {
		switch mux.Vars(r)["result"] {
		case successResult:
			_ = helpers.SetFlash("success", types.FlashMessages["account"]["subscribe"]["success"], session, r, w)
		case cancelResult:
			_ = helpers.SetFlash("error", types.FlashMessages["account"]["subscribe"]["success"], session, r, w)
		}
	}

	helpers.SetSessionHandler(helpers.AccountSubscription, session, w, r)
	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.AccountSubscription), session.Values)
	helpers.ClearFlash(session, r, w)
}

func (a *Account) InvoicesByUser(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(a.sessionStore, a.logger, w, r))

	var invoices types.Invoices
	accounts.Invoices(r.Context(), &invoices, user.UserID, a.db)

	invoices.Draw = helpers.DtDraw(r)

	amount := int64(len(invoices.Data))
	invoices.RecordsTotal = amount
	invoices.RecordsFiltered = amount

	js, err := json.Marshal(invoices)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (a *Account) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := helpers.GetSessionUser(session)

	if err := accounts.UpdateSubscription(r.Context(), user.ID, mux.Vars(r)["planID"], a.db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.SetSessionUser(user, session, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *Account) CancelSubscription(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := helpers.GetSessionUser(session)

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ok, _ := helpers.ValidateInput(helpers.ValidateCancellation{Data: r.Form.Get("cancel")}, &a.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := accounts.CancelSubscription(r.Context(), user.ID, a.db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.SetSessionUser(user, session, w, r)
	_ = helpers.SetFlash("success", types.FlashMessages["account"]["subscription"]["cancel"], session, r, w)
	w.WriteHeader(http.StatusOK)
}

func (a *Account) ReactivateSubscription(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := helpers.GetSessionUser(session)

	if err := accounts.ReactivateSubscription(r.Context(), user, mux.Vars(r)["planID"], a.db); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.SetSessionUser(user, session, w, r)
	w.WriteHeader(http.StatusOK)
}
