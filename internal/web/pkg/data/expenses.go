package data

import (
	"fmt"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/web/types"
)

func (d *Data) ExpensesCostOfGoods(userID string, expenses *types.Expenses, filter *models.Filter) {
	costOfGoods := models.LoadExpensesCostOfGoods(userID, filter, d.db)
	for _, item := range *costOfGoods {
		expenses.Data = append(expenses.Data, types.ExpensesTable{
			RowID:       item.ExpenseID,
			ProductID:   item.ProductID,
			SKU:         item.SKU,
			Description: item.Description,
			Marketplace: item.Marketplace,
			FromDate:    item.FromDate.Format("2006-01-02"),
			Amount:      item.Amount,
		})
	}

	d.expenseTotals(userID, expensesCostOfGoods, expenses)
}

func (d *Data) ExpenseCostOfGoodsForSendData(productID int64, expense *models.ExpenseCostOfGoodSendData) {
	product := models.LoadProductByID(expense.UserID, productID, d.db)
	expense.SKU = product.SKU
	expense.Marketplace = product.Marketplace
}

func (d *Data) ExpensesOther(userID string, expenses *types.Expenses, filter *models.Filter) {
	currencies := models.LoadCurrencies(d.db)

	other := models.LoadExpensesOthers(userID, filter, d.db)
	for _, item := range *other {
		expenses.Data = append(expenses.Data, types.ExpensesTable{
			RowID:       item.ExpenseID,
			CurrencyID:  d.expenseCurrency(currencies, item.CurrencyID),
			Description: item.Description,
			DateTime:    item.DateTime.Format("2006-01-02"),
			Amount:      item.Amount,
			Currency:    fmt.Sprintf("%v (%v)", item.CurrencyCode, item.CurrencySymbol),
		})
	}

	d.expenseTotals(userID, expensesOther, expenses)
}

func (d *Data) expenseCurrency(currencies *[]models.Currency, currencyID int64) string {
	for _, currency := range *currencies {
		if currency.ID == currencyID {
			return currency.CurrencyID
		}
	}

	return (*currencies)[0].CurrencyID
}

func (d *Data) expenseTotals(userID, view string, expenses *types.Expenses) {
	var total int64
	switch view {
	case expensesCostOfGoods:
		total = models.TotalExpensesCostOfGoods(userID, d.db)
	case expensesOther:
		total = models.TotalExpensesOthers(userID, d.db)
	}

	expenses.RecordsTotal = total
	expenses.RecordsFiltered = total

	if len(expenses.Data) == 0 {
		expenses.Data = []types.ExpensesTable{}
	}
}
