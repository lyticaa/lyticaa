package data

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (d *Data) ExpensesCostOfGoods(userId string, expenses *types.Expenses, filter *models.Filter) {
	costOfGoods := models.LoadExpensesCostOfGoods(userId, filter, d.db)
	for _, item := range *costOfGoods {
		expenses.Data = append(expenses.Data, types.ExpensesTable{
			SKU:         item.SKU,
			Description: item.Description,
			Marketplace: item.Marketplace,
			Cost:        item.Cost,
			FromDate:    item.FromDate,
		})
	}

	d.expenseTotals(userId, expensesCostOfGoods, expenses)
}

func (d *Data) ExpensesOther(userId string, expenses *types.Expenses, filter *models.Filter) {
	other := models.LoadExpensesOther(userId, filter, d.db)
	for _, item := range *other {
		expenses.Data = append(expenses.Data, types.ExpensesTable{
			Description: item.Description,
			Cost:        item.Cost,
			DateTime:    item.DateTime,
		})
	}

	d.expenseTotals(userId, expensesOther, expenses)
}

func (d *Data) expenseTotals(userId, view string, expenses *types.Expenses) {
	var total int64
	switch view {
	case expensesCostOfGoods:
		total = models.TotalExpensesCostOfGoods(userId, d.db)
	case expensesOther:
		total = models.TotalExpensesOther(userId, d.db)
	}

	expenses.RecordsTotal = total
	expenses.RecordsFiltered = total

	if len(expenses.Data) == 0 {
		expenses.Data = []types.ExpensesTable{}
	}
}
