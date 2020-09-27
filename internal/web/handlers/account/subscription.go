package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/lyticaa/lyticaa-app/internal/web/helpers"
	"gitlab.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
	"golang.org/x/text/currency"
)

var (
	invoiceStatusMap = map[string]string{
		"draft":         "badge-info",
		"open":          "badge-warning",
		"paid":          "badge-success",
		"void":          "badge-info",
		"uncollectible": "badge-danger",
	}
)

type Cancellation struct {
	Data string `validate:"required,eq=CANCEL"`
}

func (a *Account) Subscription(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	helpers.RenderTemplate(w, helpers.TemplateList(helpers.AccountSubscription), session.Values)
	helpers.ClearFlash(session, r, w)
}

func (a *Account) InvoicesByUser(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(a.sessionStore, a.logger, w, r))

	js, err := json.Marshal(a.paintInvoices(user.StripeUserId.String, r))
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

	planId := mux.Vars(r)["planId"]
	if user.StripePlanId.String == planId {
		a.logger.Error().Msgf("user %v already on plan %v", user.UserId, user.StripePlanId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := a.stripe.ChangePlan(user.StripeSubscriptionId.String, planId); err != nil {
		a.logger.Error().Err(err).Msg("unable to change the plan")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var plan sql.NullString
	if err := plan.Scan(planId); err != nil {
		a.logger.Error().Err(err).Msg("unable to assign stripe plan id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.StripePlanId = plan
	_ = user.Save(a.db)

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

	err := a.stripe.CancelSubscription(user.StripeSubscriptionId.String)
	if err != nil {
		a.logger.Error().Err(err).Msg("unable to cancel subscription")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.StripeSubscriptionId = sql.NullString{}
	user.StripePlanId = sql.NullString{}
	_ = user.Save(a.db)

	helpers.SetSessionUser(user, session, w, r)

	_ = helpers.SetFlash("success", types.FlashMessages["account"]["subscription"]["cancel"], session, r, w)
	w.WriteHeader(http.StatusOK)
}

func (a *Account) Subscribe(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := helpers.GetSessionUser(session)

	planId := mux.Vars(r)["planId"]
	if user.StripePlanId.String == planId {
		a.logger.Error().Msgf("user %v already on plan %v", user.UserId, user.StripePlanId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	s, err := a.stripe.CreateSubscription(user.StripeUserId.String, planId)
	if err != nil {
		a.logger.Error().Err(err).Msg("unable to subscribe")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var subscription sql.NullString
	if err := subscription.Scan(s.ID); err != nil {
		a.logger.Error().Err(err).Msg("unable to assign stripe plan id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var plan sql.NullString
	if err := plan.Scan(planId); err != nil {
		a.logger.Error().Err(err).Msg("unable to assign stripe plan id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.StripeSubscriptionId = subscription
	user.StripePlanId = plan
	_ = user.Save(a.db)

	helpers.SetSessionUser(user, session, w, r)

	w.WriteHeader(http.StatusOK)
}

func (a *Account) paintInvoices(stripeUserId string, r *http.Request) types.Invoices {
	var byUser types.Invoices

	invoices := a.stripe.InvoicesByUser(stripeUserId)
	for _, invoice := range *invoices {
		unit, _ := currency.ParseISO(string(invoice.Currency))

		t := types.InvoiceTable{
			Number:      invoice.Number,
			Date:        invoice.Date.Format("2006-01-02"),
			Amount:      fmt.Sprintf("%v %v", currency.Symbol(unit), invoice.Amount),
			Status:      strings.ToUpper(string(invoice.Status)),
			StatusClass: a.invoiceClass(string(invoice.Status)),
			PDF:         invoice.PDF,
		}

		byUser.Data = append(byUser.Data, t)
	}

	if len(byUser.Data) == 0 {
		byUser.Data = []types.InvoiceTable{}
	}

	byUser.Draw = helpers.DtDraw(r)

	amount := int64(len(byUser.Data))
	byUser.RecordsTotal = amount
	byUser.RecordsFiltered = amount

	return byUser
}

func (a *Account) invoiceClass(status string) string {
	return invoiceStatusMap[status]
}
