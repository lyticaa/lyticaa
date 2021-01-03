package cost_of_goods

import (
	"context"
	"time"

	"github.com/lyticaa/lyticaa-app/internal/app/types"
	"github.com/lyticaa/lyticaa-app/internal/models"

	"github.com/jmoiron/sqlx"
)

func ExpensesCostOfGood(ctx context.Context, expenseID string, db *sqlx.DB) models.ExpensesCostOfGoodModel {
	expensesCostOfGoodModel := &models.ExpensesCostOfGoodModel{
		ExpenseID: expenseID,
	}

	expenseCostOfGood := expensesCostOfGoodModel.FetchOne(ctx, db).(models.ExpensesCostOfGoodModel)

	return expenseCostOfGood
}

func ExpensesCostOfGoods(ctx context.Context, expenses *types.Expenses, filter *models.Filter, userID string, db *sqlx.DB) {
	var costOfGoodModel models.ExpensesCostOfGoodModel

	data := make(map[string]interface{})
	data["UserID"] = userID

	costOfGoods := costOfGoodModel.FetchAll(ctx, data, filter, db).([]models.ExpensesCostOfGoodModel)
	for _, costOfGood := range costOfGoods {
		table := types.ExpensesTable{
			RowID:       costOfGood.ExpenseID,
			ProductID:   costOfGood.ProductID,
			SKU:         costOfGood.SKU,
			Description: costOfGood.Description,
			Marketplace: costOfGood.Marketplace,
			FromDate:    costOfGood.FromDate.Format("2006-01-02"),
			Amount:      costOfGood.Amount,
		}

		expenses.Data = append(expenses.Data, table)
	}

	expenses.RecordsTotal = costOfGoodModel.Count(ctx, data, db)
	expenses.RecordsFiltered = expenses.RecordsTotal
}

func CreateExpensesCostOfGood(ctx context.Context, productID int64, description string, amount float64, fromDate time.Time, db *sqlx.DB) error {
	expensesCostOfGoodModel := &models.ExpensesCostOfGoodModel{
		ProductID:   productID,
		Description: description,
		Amount:      amount,
		FromDate:    fromDate,
	}

	if err := expensesCostOfGoodModel.Create(ctx, db); err != nil {
		return err
	}

	return nil
}

func UpdateExpensesCostOfGood(ctx context.Context, expense *models.ExpensesCostOfGoodModel, description string, amount float64, fromDate time.Time, db *sqlx.DB) error {
	expense.Description = description
	expense.Amount = amount
	expense.FromDate = fromDate

	if err := expense.Update(ctx, db); err != nil {
		return err
	}

	return nil
}

func DeleteExpensesCostOfGood(ctx context.Context, expense *models.ExpensesCostOfGoodModel, db *sqlx.DB) error {
	if err := expense.Delete(ctx, db); err != nil {
		return err
	}

	return nil
}

func Products(ctx context.Context, userID int64, filter *models.Filter, db *sqlx.DB) []types.Product {
	productModel := &models.ProductModel{
		UserID: userID,
	}

	var productList []types.Product

	products := productModel.FetchAll(ctx, nil, filter, db).([]models.ProductModel)
	for _, product := range products {
		productList = append(productList, types.Product{
			ProductID:   product.ProductID,
			SKU:         product.SKU,
			Marketplace: product.Marketplace,
			Description: product.Description,
		})
	}

	return productList
}

func ProductOwner(ctx context.Context, userID int64, productID string, db *sqlx.DB) (bool, *models.ProductModel) {
	productModel := &models.ProductModel{
		ProductID: productID,
		UserID:    userID,
	}

	product := productModel.FetchBy(ctx, db).(*models.ProductModel)
	if product.UserID == userID {
		return true, product
	}

	return false, nil
}
