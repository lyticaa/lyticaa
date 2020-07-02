package expenses

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (e *Expenses) CostOfGoods(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"partials/filters/_filters",
		"partials/filters/_upload",
		"expenses/cost_of_goods",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (e *Expenses) CostOfGoodsByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)
	user := session.Values["User"].(models.User)

	var byDate types.Expenses
	byDate.Draw = helpers.DtDraw(r)

	e.data.ExpensesCostOfGoods(user.UserId, &byDate, helpers.BuildFilter(r))
	js, err := json.Marshal(byDate)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (e *Expenses) NewCostOfGood(w http.ResponseWriter, r *http.Request) {

}

func (e *Expenses) Products(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(e.sessionStore, e.logger, w, r)
	user := session.Values["User"].(models.User)

	js, err := json.Marshal(e.paintProducts(user.UserId))
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
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
