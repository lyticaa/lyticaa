package expenses

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
)

type ValidateOther struct {
	CurrencyID  string    `validate:"required,uuid4"`
	Description string    `validate:"required,min=3"`
	DateTime    time.Time `validate:"required"`
	Amount      float64   `validate:"required,gt=0"`
}

func (e *Expenses) Other(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)
	helpers.RenderTemplate(w, helpers.TemplateList(helpers.ExpensesOther), session.Values)
}

func (e *Expenses) OtherByUser(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	var byDate types.Expenses
	byDate.Draw = helpers.DtDraw(r)

	e.data.ExpensesOther(user.UserID, &byDate, helpers.BuildFilter(r))
	js, err := json.Marshal(byDate)
	if err != nil {
		e.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (e *Expenses) NewOther(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	currencyID := r.FormValue("currency")
	description := r.FormValue("description")
	dateTime, _ := time.Parse("2006-01-02", r.FormValue("dateTime"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	ok, res := e.validateOther(currencyID, description, dateTime, amount)
	if !ok {
		js, err := json.Marshal(res)
		if err != nil {
			e.logger.Error().Err(err).Msg("failed to marshal data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write(js)
		return
	}

	ok, currency := e.isValidCurrency(currencyID)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := models.CreateExpensesOther(user.UserID, currency.ID, description, amount, dateTime, e.db); err != nil {
		e.logger.Error().Err(err).Msg("failed to create the expense")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *Expenses) EditOther(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	params := mux.Vars(r)
	expenseID := params["expense"]

	ok, _ := helpers.ValidateInput(ValidateExpense{ExpenseID: expenseID}, &e.logger)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expense := models.LoadExpensesOther(expenseID, e.db)
	if expense.UserID != user.UserID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	currencyID := r.FormValue("currency")
	description := r.FormValue("description")
	dateTime, _ := time.Parse("2006-01-02", r.FormValue("dateTime"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	ok, res := e.validateOther(currencyID, description, dateTime, amount)
	if !ok {
		js, err := json.Marshal(res)
		if err != nil {
			e.logger.Error().Err(err).Msg("failed to marshal data")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write(js)
		return
	}

	ok, currency := e.isValidCurrency(currencyID)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	expense.CurrencyID = currency.ID
	expense.Description = description
	expense.DateTime = dateTime
	expense.Amount = amount

	if err := expense.Save(e.db); err != nil {
		e.logger.Error().Err(err).Msg("failed to update the expense")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *Expenses) DeleteOther(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	params := mux.Vars(r)
	expenseID := params["expense"]

	ok, _ := helpers.ValidateInput(ValidateExpense{ExpenseID: expenseID}, &e.logger)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expense := models.LoadExpensesOther(expenseID, e.db)
	if expense.UserID != user.UserID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := expense.Delete(e.db); err != nil {
		e.logger.Error().Err(err).Msgf("failed to delete the expense %v", expenseID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
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
			CurrencyID: currency.CurrencyID,
			Code:       currency.Code,
			Symbol:     currency.Symbol,
		})
	}

	return &currencyList
}

func (e *Expenses) isValidCurrency(currencyID string) (bool, *models.Currency) {
	isValidCurrency := false

	currency := models.LoadCurrency(currencyID, e.db)
	if currency.CurrencyID == currencyID {
		isValidCurrency = true
	}

	return isValidCurrency, currency
}

func (e *Expenses) validateOther(currencyID, description string, dateTime time.Time, amount float64) (bool, map[string]string) {
	return helpers.ValidateInput(
		ValidateOther{
			CurrencyID:  currencyID,
			Description: description,
			DateTime:    dateTime,
			Amount:      amount,
		},
		&e.logger)
}
