package expenses

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"

	"github.com/gorilla/mux"
)

type ValidateCostOfGood struct {
	ProductId   string    `validate:"required,uuid4"`
	Description string    `validate:"required,min=3"`
	FromDate    time.Time `validate:"required"`
	Amount      float64   `validate:"required,gt=0"`
}

func (e *Expenses) CostOfGoods(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_account",
		"partials/nav/account/_main",
		"partials/admin/_impersonate",
		"partials/filters/_filters",
		"partials/filters/_import",
		"partials/expenses/cost_of_goods/_form",
		"expenses/cost_of_goods",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (e *Expenses) CostOfGoodsByUser(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	var byUser types.Expenses
	byUser.Draw = helpers.DtDraw(r)

	e.data.ExpensesCostOfGoods(user.UserId, &byUser, helpers.BuildFilter(r))
	js, err := json.Marshal(byUser)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (e *Expenses) NewCostOfGood(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	productId := r.FormValue("product")
	ok, product := e.isProductOwner(user.UserId, productId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	fromDate, _ := time.Parse("2006-01-02", r.FormValue("fromDate"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	ok, res := e.validateCostOfGood(productId, description, fromDate, amount)
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

	if err := models.CreateExpensesCostOfGood(product.Id, description, amount, fromDate, e.db); err != nil {
		e.logger.Error().Err(err).Msg("failed to create the expense")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *Expenses) Products(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	js, err := json.Marshal(e.paintProducts(user.UserId))
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (e *Expenses) EditCostOfGood(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	params := mux.Vars(r)
	expenseId := params["expense"]

	ok, _ := helpers.ValidateInput(ValidateExpense{ExpenseId: expenseId}, &e.logger)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expense := models.LoadExpensesCostOfGood(expenseId, e.db)
	ok, _ = e.isProductOwner(user.UserId, expense.ProductId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	fromDate, _ := time.Parse("2006-01-02", r.FormValue("fromDate"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	ok, res := e.validateCostOfGood(expense.ProductId, description, fromDate, amount)
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

	expense.Description = description
	expense.FromDate = fromDate
	expense.Amount = amount

	if err := expense.Save(e.db); err != nil {
		e.logger.Error().Err(err).Msg("failed to update the expense")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *Expenses) DeleteCostOfGood(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	params := mux.Vars(r)
	expenseId := params["expense"]

	ok, _ := helpers.ValidateInput(ValidateExpense{ExpenseId: expenseId}, &e.logger)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expense := models.LoadExpensesCostOfGood(expenseId, e.db)
	ok, _ = e.isProductOwner(user.UserId, expense.ProductId)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := expense.Delete(e.db); err != nil {
		e.logger.Error().Err(err).Msgf("failed to delete the expense %v", expenseId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *Expenses) paintProducts(userId string) *[]types.Product {
	var productList []types.Product

	products := models.LoadProducts(userId, e.db)
	for _, product := range *products {
		productList = append(productList, types.Product{
			ProductId:   product.ProductId,
			SKU:         product.SKU,
			Marketplace: product.Marketplace,
			Description: product.Description,
		})
	}

	if len(productList) == 0 {
		productList = []types.Product{}
	}

	return &productList
}

func (e *Expenses) isProductOwner(userId, productId string) (bool, *models.Product) {
	isProductOwner := false

	product := models.LoadProduct(userId, productId, e.db)
	if product.UserId == userId {
		isProductOwner = true
	}

	return isProductOwner, product
}

func (e *Expenses) validateCostOfGood(productId, description string, fromDate time.Time, amount float64) (bool, map[string]string) {
	return helpers.ValidateInput(
		ValidateCostOfGood{
			ProductId:   productId,
			Description: description,
			FromDate:    fromDate,
			Amount:      amount,
		},
		&e.logger)
}
