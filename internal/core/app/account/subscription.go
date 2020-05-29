package account

import (
	"encoding/json"
	"fmt"
	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/core/payments"
	"gitlab.com/getlytica/lytica-app/internal/models"
	"golang.org/x/text/currency"
	"net/http"
	"strings"
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

func (a *Account) Subscription(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)

	t := []string{
		helpers.NavForSession(helpers.IsSubscribed(a.sessionStore, a.logger, w, r)),
		"account/subscription",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (a *Account) InvoicesByUser(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := session.Values["User"].(models.User)

	var byUser types.Invoices

	invoices := payments.InvoicesByUser(user.StripeUserId.String)
	for _, invoice := range *invoices {
		unit, _ := currency.ParseISO(string(invoice.Currency))

		t := types.InvoiceTable{
			Number:      invoice.Number,
			Date:        fmt.Sprintf("%s", invoice.Date.Format("2006-01-02")),
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

	js, err := json.Marshal(byUser)
	if err != nil {
		a.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (a *Account) invoiceClass(status string) string {
	return invoiceStatusMap[status]
}
