package expenses

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/expenses/other"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (e *Expenses) Other(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)
	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.ExpensesOther), session.Values)
}

func (e *Expenses) OtherByUser(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	var expenses types.Expenses
	expenses.Draw = helpers.DtDraw(r)

	other.ExpensesOthers(r.Context(), &expenses, helpers.BuildFilter(r), user.UserID, e.db)
	js, err := json.Marshal(expenses)
	if err != nil {
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

	ok, result := helpers.ValidateInput(
		helpers.ValidateExpensesOther{
			CurrencyID:  currencyID,
			Description: description,
			DateTime:    dateTime,
			Amount:      amount,
		},
		&e.logger)

	if !ok {
		js, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write(js)
		return
	}

	ok, currency := other.SupportedCurrency(r.Context(), currencyID, e.db)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := other.CreateExpensesOther(r.Context(), user.ID, currency.ID, description, amount, dateTime, e.db); err != nil {
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

	expense := other.FetchExpensesOther(r.Context(), expenseID, e.db)
	if expense.UserID != user.ID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	currencyID := r.FormValue("currency")
	description := r.FormValue("description")
	dateTime, _ := time.Parse("2006-01-02", r.FormValue("dateTime"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	ok, result := helpers.ValidateInput(
		helpers.ValidateExpensesOther{
			CurrencyID:  currencyID,
			Description: description,
			DateTime:    dateTime,
			Amount:      amount,
		},
		&e.logger)

	if !ok {
		js, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write(js)
		return
	}

	ok, currency := other.SupportedCurrency(r.Context(), currencyID, e.db)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if err := other.UpdateExpensesCostOfGood(r.Context(), &expense, currency.ID, description, amount, dateTime, e.db); err != nil {
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

	expense := other.FetchExpensesOther(r.Context(), expenseID, e.db)
	if expense.UserID != user.ID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := other.DeleteExpensesOther(r.Context(), &expense, e.db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *Expenses) Currencies(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(other.Currencies(r.Context(), e.db))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
