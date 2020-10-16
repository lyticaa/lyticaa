package data

import (
	"fmt"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/web/types"
)

func (d *Data) ExpensesCostOfGoods(userId string, expenses *types.Expenses, filter *models.Filter) {
	costOfGoods := models.LoadExpensesCostOfGoods(userId, filter, d.db)
	for _, item := range *costOfGoods {
		expenses.Data = append(expenses.Data, types.ExpensesTable{
			RowId:       item.ExpenseId,
			ProductId:   item.ProductId,
			SKU:         item.SKU,
			Description: item.Description,
			Marketplace: item.Marketplace,
			FromDate:    item.FromDate.Format("2006-01-02"),
			Amount:      item.Amount,
		})
	}

	d.expenseTotals(userId, expensesCostOfGoods, expenses)
}

func (d *Data) ExpensesOther(userId string, expenses *types.Expenses, filter *models.Filter) {
	currencies := models.LoadCurrencies(d.db)

	other := models.LoadExpensesOthers(userId, filter, d.db)
	for _, item := range *other {
		expenses.Data = append(expenses.Data, types.ExpensesTable{
			RowId:       item.ExpenseId,
			CurrencyId:  d.expenseCurrency(currencies, item.CurrencyId),
			Description: item.Description,
			DateTime:    item.DateTime.Format("2006-01-02"),
			Amount:      item.Amount,
			Currency:    fmt.Sprintf("%v (%v)", item.CurrencyCode, item.CurrencySymbol),
		})
	}

	d.expenseTotals(userId, expensesOther, expenses)
}

func (d *Data) expenseCurrency(currencies *[]models.Currency, currencyId int64) string {
	for _, currency := range *currencies {
		if currency.Id == currencyId {
			return currency.CurrencyId
		}
	}

	return (*currencies)[0].CurrencyId
}

func (d *Data) expenseTotals(userId, view string, expenses *types.Expenses) {
	var total int64
	switch view {
	case expensesCostOfGoods:
		total = models.TotalExpensesCostOfGoods(userId, d.db)
	case expensesOther:
		total = models.TotalExpensesOthers(userId, d.db)
	}

	expenses.RecordsTotal = total
	expenses.RecordsFiltered = total

	if len(expenses.Data) == 0 {
		expenses.Data = []types.ExpensesTable{}
	}
}
