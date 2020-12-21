package account

import (
	"encoding/json"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/accounts"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
)

type Cancellation struct {
	Data string `validate:"required,eq=CANCEL"`
}

func (a *Account) Subscription(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.AccountSubscription), session.Values)
	helpers.ClearFlash(session, r, w)
}

func (a *Account) InvoicesByUser(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(a.sessionStore, a.logger, w, r))

	var invoices types.Invoices
	accounts.Invoices(r.Context(), &invoices, user.ID, a.db)

	invoices.Draw = helpers.DtDraw(r)

	amount := int64(len(invoices.Data))
	invoices.RecordsTotal = amount
	invoices.RecordsFiltered = amount

	js, err := json.Marshal(invoices)
	if err != nil {
		a.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (a *Account) ChangePlan(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := helpers.GetSessionUser(session)

	if err := accounts.ChangePlan(r.Context(), user.ID, mux.Vars(r)["planID"], a.db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.SetSessionUser(user, session, w, r)
	w.WriteHeader(http.StatusOK)
}

func (a *Account) CancelSubscription(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := helpers.GetSessionUser(session)

	cancel := r.FormValue("cancel")
	validate := Cancellation{Data: cancel}

	ok, _ := helpers.ValidateInput(validate, &a.logger)
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

func (a *Account) Subscribe(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := helpers.GetSessionUser(session)

	if err := accounts.Subscribe(r.Context(), user.ID, mux.Vars(r)["planID"], a.db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.SetSessionUser(user, session, w, r)
	w.WriteHeader(http.StatusOK)
}
