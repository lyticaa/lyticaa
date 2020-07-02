package expenses

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (e *Expenses) Other(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"partials/filters/_filters",
		"partials/filters/_upload",
		"expenses/other",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (e *Expenses) OtherByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)
	user := session.Values["User"].(models.User)

	var byDate types.Expenses
	byDate.Draw = helpers.DtDraw(r)

	e.data.ExpensesOther(user.UserId, &byDate, helpers.BuildFilter(r))
	js, err := json.Marshal(byDate)
	if err != nil {
		e.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (e *Expenses) Currencies(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(e.paintCurrencies())
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (e *Expenses) paintCurrencies() *[]types.Currency {
	var currencyList []types.Currency

	currencies := models.LoadCurrencies(e.db)
	for _, currency := range *currencies {
		currencyList = append(currencyList, types.Currency{
			CurrencyId: currency.CurrencyId,
			Code:       currency.Code,
			Symbol:     currency.Symbol,
		})
	}

	return &currencyList
}
