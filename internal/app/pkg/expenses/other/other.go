package other

import (
	"context"
	"time"

	"github.com/lyticaa/lyticaa-app/internal/app/types"
	"github.com/lyticaa/lyticaa-app/internal/models"

	"github.com/jmoiron/sqlx"
)

func ExpensesOther(ctx context.Context, expenseID string, db *sqlx.DB) models.ExpensesOtherModel {
	expensesOtherModel := &models.ExpensesOtherModel{
		ExpenseID: expenseID,
	}

	expenseOther := expensesOtherModel.FetchOne(ctx, db).(models.ExpensesOtherModel)

	return expenseOther
}

func ExpensesOthers(ctx context.Context, expenses *types.Expenses, filter *models.Filter, userID string, db *sqlx.DB) {
	var otherModel models.ExpensesOtherModel

	data := make(map[string]interface{})
	data["UserID"] = userID

	others := otherModel.FetchAll(ctx, data, filter, db).([]models.ExpensesOtherModel)
	for _, other := range others {
		table := types.ExpensesTable{
			RowID:       other.ExpenseID,
			Description: other.Description,
			DateTime:    other.DateTime.Format("2006-01-02"),
			Amount:      other.Amount,
			Currency:    other.CurrencyCode,
		}

		expenses.Data = append(expenses.Data, table)
	}

	expenses.RecordsTotal = otherModel.Count(ctx, data, db)
	expenses.RecordsFiltered = expenses.RecordsTotal
}

func CreateExpensesOther(ctx context.Context, userID int64, currencyID int64, description string, amount float64, dateTime time.Time, db *sqlx.DB) error {
	otherModel := &models.ExpensesOtherModel{
		UserID:      userID,
		CurrencyID:  currencyID,
		Description: description,
		Amount:      amount,
		DateTime:    dateTime,
	}

	if err := otherModel.Create(ctx, db); err != nil {
		return err
	}

	return nil
}

func UpdateExpensesCostOfGood(ctx context.Context, expense *models.ExpensesOtherModel, currencyID int64, description string, amount float64, dateTime time.Time, db *sqlx.DB) error {
	expense.CurrencyID = currencyID
	expense.Description = description
	expense.Amount = amount
	expense.DateTime = dateTime

	if err := expense.Update(ctx, db); err != nil {
		return err
	}

	return nil
}

func DeleteExpensesOther(ctx context.Context, expense *models.ExpensesOtherModel, db *sqlx.DB) error {
	if err := expense.Delete(ctx, db); err != nil {
		return err
	}

	return nil
}

func Currencies(ctx context.Context, db *sqlx.DB) []types.Currency {
	var currencyModel models.CurrencyModel
	var currencies []types.Currency

	currencyList := currencyModel.FetchAll(ctx, nil, nil, db).([]models.CurrencyModel)
	for _, currency := range currencyList {
		currencies = append(currencies, types.Currency{
			CurrencyID: currency.CurrencyID,
			Code:       currency.Code,
		})
	}

	return currencies
}

func SupportedCurrency(ctx context.Context, currencyID string, db *sqlx.DB) (bool, *models.CurrencyModel) {
	currencyModel := &models.CurrencyModel{
		CurrencyID: currencyID,
	}

	currency := currencyModel.FetchOne(ctx, db).(models.CurrencyModel)
	if currency.CurrencyID == currencyID {
		return true, &currency
	}

	return false, nil
}
