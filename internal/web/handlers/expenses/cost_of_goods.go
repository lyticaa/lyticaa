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

const amazonExpenses = "/Expenses/Amazon"

type ValidateCostOfGood struct {
	ProductID   string    `validate:"required,uuid4"`
	Description string    `validate:"required,min=3"`
	FromDate    time.Time `validate:"required"`
	Amount      float64   `validate:"required,gt=0"`
}

func (e *Expenses) CostOfGoods(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)
	helpers.RenderTemplate(w, helpers.TemplateList(helpers.ExpensesCostOfGoods), session.Values)
}

func (e *Expenses) CostOfGoodsByUser(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	var byUser types.Expenses
	byUser.Draw = helpers.DtDraw(r)

	e.data.ExpensesCostOfGoods(user.UserID, &byUser, helpers.BuildFilter(r))
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

	productID := r.FormValue("product")
	ok, product := e.isProductOwner(user.UserID, productID)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	fromDate, _ := time.Parse("2006-01-02", r.FormValue("fromDate"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	ok, res := e.validateCostOfGood(productID, description, fromDate, amount)
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

	expenseID, err := models.CreateExpensesCostOfGood(product.ID, description, amount, fromDate, e.db)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to create the expense")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	expense := models.ExpenseCostOfGoodSendData{UserID: user.UserID, Description: description, ExpenseID: expenseID, Amount: amount, DateTime: fromDate}

	e.data.ExpenseCostOfGoodsForSendData(product.ID, &expense)

	js, err := json.Marshal(expense)
	if err != nil {
		e.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := sendMessage(user.UserID, amazonExpenses, string(js)); err != nil {
		e.logger.Error().Err(err).Msgf("failed to send message into queue costOfGood for the user %v", user.UserID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *Expenses) Products(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(e.sessionStore, e.logger, w, r))

	js, err := json.Marshal(e.paintProducts(user.UserID))
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
	expenseID := params["expense"]

	ok, _ := helpers.ValidateInput(ValidateExpense{ExpenseID: expenseID}, &e.logger)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expense := models.LoadExpensesCostOfGood(expenseID, e.db)
	ok, _ = e.isProductOwner(user.UserID, expense.ProductID)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	fromDate, _ := time.Parse("2006-01-02", r.FormValue("fromDate"))
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)

	ok, res := e.validateCostOfGood(expense.ProductID, description, fromDate, amount)
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
	expenseID := params["expense"]

	ok, _ := helpers.ValidateInput(ValidateExpense{ExpenseID: expenseID}, &e.logger)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expense := models.LoadExpensesCostOfGood(expenseID, e.db)
	ok, _ = e.isProductOwner(user.UserID, expense.ProductID)
	if !ok {
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

func (e *Expenses) paintProducts(userID string) *[]types.Product {
	var productList []types.Product

	products := models.LoadProducts(userID, e.db)
	for _, product := range *products {
		productList = append(productList, types.Product{
			ProductID:   product.ProductID,
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

func (e *Expenses) isProductOwner(userID, productID string) (bool, *models.Product) {
	isProductOwner := false

	product := models.LoadProduct(userID, productID, e.db)
	if product.UserID == userID {
		isProductOwner = true
	}

	return isProductOwner, product
}

func (e *Expenses) validateCostOfGood(productID, description string, fromDate time.Time, amount float64) (bool, map[string]string) {
	return helpers.ValidateInput(
		ValidateCostOfGood{
			ProductID:   productID,
			Description: description,
			FromDate:    fromDate,
			Amount:      amount,
		},
		&e.logger)
}
