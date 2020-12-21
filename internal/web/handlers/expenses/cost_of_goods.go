package expenses

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	cog "github.com/lyticaa/lyticaa-app/internal/web/pkg/expenses/cost_of_goods"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (e *Expenses) CostOfGoods(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)
	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.ExpensesCostOfGoods), session.Values)
}

func (e *Expenses) CostOfGoodsByUser(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	var expenses types.Expenses
	expenses.Draw = helpers.DtDraw(r)

	cog.ExpensesCostOfGoods(r.Context(), &expenses, helpers.BuildFilter(r), user.UserID, e.db)
	js, err := json.Marshal(expenses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (e *Expenses) NewCostOfGood(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	productID := r.FormValue("product")
	ok, product := cog.ProductOwner(r.Context(), user.ID, productID, e.db)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	fromDate, _ := time.Parse("2006-01-02", r.FormValue("fromDate"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	ok, result := helpers.ValidateInput(
		helpers.ValidateExpensesCostOfGood{
			ProductID:   productID,
			Description: description,
			FromDate:    fromDate,
			Amount:      amount,
		},
		&e.logger,
	)

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

	if err := cog.CreateExpensesCostOfGood(r.Context(), product.ID, description, amount, fromDate, e.db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *Expenses) Products(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	products := cog.Products(r.Context(), user.ID, nil, e.db)
	js, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (e *Expenses) EditCostOfGood(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	params := mux.Vars(r)
	expenseID := params["expense"]
	productID := params["product"]

	ok, _ := helpers.ValidateInput(ValidateExpense{ExpenseID: expenseID}, &e.logger)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expense := cog.FetchExpensesCostOfGood(r.Context(), expenseID, e.db)
	ok, product := cog.ProductOwner(r.Context(), user.ID, productID, e.db)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if product.ID != expense.ProductID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	fromDate, _ := time.Parse("2006-01-02", r.FormValue("fromDate"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	ok, result := helpers.ValidateInput(
		helpers.ValidateExpensesCostOfGood{
			ProductID:   product.ProductID,
			Description: description,
			FromDate:    fromDate,
			Amount:      amount,
		},
		&e.logger,
	)

	if !ok {
		js, err := json.Marshal(result)
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

	if err := cog.UpdateExpensesCostOfGood(r.Context(), &expense, description, amount, fromDate, e.db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *Expenses) DeleteCostOfGood(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	params := mux.Vars(r)
	expenseID := params["expense"]
	productID := params["product"]

	ok, _ := helpers.ValidateInput(ValidateExpense{ExpenseID: expenseID}, &e.logger)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expense := cog.FetchExpensesCostOfGood(r.Context(), expenseID, e.db)
	ok, product := cog.ProductOwner(r.Context(), user.ID, productID, e.db)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if product.ID != expense.ProductID {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := cog.DeleteExpensesCostOfGood(r.Context(), &expense, e.db); err != nil {
		e.logger.Error().Err(err).Msgf("failed to delete the expense %v", expenseID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
