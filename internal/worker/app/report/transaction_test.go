package report

import (
	"bytes"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	transactionType    = "Order"
	marketplace        = "amazon.com"
	fulfillment        = "Amazon"
	taxCollectionModel = "MarketplaceFacilitator"
)

func TestGetTransactionTypes(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, transactionType, time.Now(), time.Now())
	mock.ExpectQuery("^SELECT (.+) FROM transaction_types").WillReturnRows(rows)

	transactionTypes := r.getTransactionTypes()
	if len(transactionTypes) == 0 {
		t.Errorf("no transaction types found!")
	}

	if len(transactionTypes) > 0 {
		if transactionTypes[0].Name != transactionType {
			t.Errorf("transaction type is invalid")
		}
	}
}

func TestGetTransactionTypeIdByName(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, transactionType, time.Now(), time.Now())
	mock.ExpectQuery("^SELECT (.+) FROM transaction_types").WillReturnRows(rows)

	transactionTypes := r.getTransactionTypes()
	_, ok := r.getTransactionTypeIdByName(transactionType, transactionTypes)
	if !ok {
		t.Errorf("transaction type not found")
	}
}

func TestGetMarketplaces(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, marketplace, time.Now(), time.Now())
	mock.ExpectQuery("^SELECT (.+) FROM marketplaces").WillReturnRows(rows)

	marketplaces := r.getMarketplaces()
	if len(marketplaces) == 0 {
		t.Errorf("no marketplaces found!")
	}

	if len(marketplaces) > 0 {
		if marketplaces[0].Name != marketplace {
			t.Errorf("marketplace is invalid")
		}
	}
}

func TestGetMarketplaceIdByName(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, marketplace, time.Now(), time.Now())
	mock.ExpectQuery("^SELECT (.+) FROM marketplaces").WillReturnRows(rows)

	marketplaces := r.getMarketplaces()
	_, ok := r.getMarketplaceIdByName(marketplace, marketplaces)
	if !ok {
		t.Errorf("marketplace not found")
	}
}

func TestGetFulfillments(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, fulfillment, time.Now(), time.Now())
	mock.ExpectQuery("^SELECT (.+) FROM fulfillments").WillReturnRows(rows)

	fulfillments := r.getFulfillments()
	if len(fulfillments) == 0 {
		t.Errorf("no fulfillments found!")
	}

	if len(fulfillments) > 0 {
		if fulfillments[0].Name != fulfillment {
			t.Errorf("fulfillment is invalid")
		}
	}
}

func TestGetFulfillmentIdByName(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, fulfillment, time.Now(), time.Now())
	mock.ExpectQuery("^SELECT (.+) FROM fulfillments").WillReturnRows(rows)

	fulfillments := r.getFulfillments()
	_, ok := r.getFulfillmentIdByName(fulfillment, fulfillments)
	if !ok {
		t.Errorf("fulfilment not found")
	}
}

func TestGetTaxCollectionModels(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, taxCollectionModel, time.Now(), time.Now())
	mock.ExpectQuery("^SELECT (.+) FROM tax_collection_models").WillReturnRows(rows)

	taxCollectionModels := r.getTaxCollectionModels()
	if len(taxCollectionModels) == 0 {
		t.Errorf("no tax collection models found!")
	}

	if len(taxCollectionModels) > 0 {
		if taxCollectionModels[0].Name != taxCollectionModel {
			t.Errorf("tax collection model is invalid")
		}
	}
}

func TestGetTaxCollectionModelIdByName(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, taxCollectionModel, time.Now(), time.Now())
	mock.ExpectQuery("^SELECT (.+) FROM tax_collection_models").WillReturnRows(rows)

	taxCollectionModels := r.getTaxCollectionModels()
	_, ok := r.getTaxCollectionModelIdByName(taxCollectionModel, taxCollectionModels)
	if !ok {
		t.Errorf("tax collection model not found")
	}
}

func TestFormatTransactions(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	body := []byte(transactionReport)
	content := r.mapCsv(transactionReportFile, bytes.NewBuffer(body))

	user := sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@example.com", time.Now(), time.Now())
	transactionTypeRows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, transactionType, time.Now(), time.Now())
	marketplaceRows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, marketplace, time.Now(), time.Now())
	fulfillmentRows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, fulfillment, time.Now(), time.Now())
	taxCollectionModelRows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, taxCollectionModel, time.Now(), time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	mock.ExpectQuery("^SELECT (.+) FROM transaction_types").WillReturnRows(transactionTypeRows)
	mock.ExpectQuery("^SELECT (.+) FROM marketplaces").WillReturnRows(marketplaceRows)
	mock.ExpectQuery("^SELECT (.+) FROM fulfillments").WillReturnRows(fulfillmentRows)
	mock.ExpectQuery("^SELECT (.+) FROM tax_collection_models").WillReturnRows(taxCollectionModelRows)

	formatted := r.formatTransactions(content, userId)

	if len(formatted) == 0 {
		t.Errorf("no formatted transactions found!")
	}

	if len(formatted) > 0 {
		if formatted[0].User.Id != 1 {
			t.Errorf("invalid user Id")
		}

		if formatted[0].SettlementIdx != 1 {
			t.Errorf("invalid settlement index")
		}

		if formatted[0].SettlementId != settlementId {
			t.Errorf("invalid settlement Id")
		}

		if formatted[0].TransactionType.Id != 1 {
			t.Errorf("invalid transaction type")
		}

		if formatted[0].Sku != sku {
			t.Errorf("invalid sku")
		}

		if formatted[0].Quantity != quantity {
			t.Errorf("invalid quantity")
		}

		if formatted[0].Marketplace.Id != 1 {
			t.Errorf("invalid marketplace")
		}

		if formatted[0].Fulfillment.Id != 1 {
			t.Errorf("invalid fulfillment")
		}
	}
}

func TestSaveTransaction(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	body := []byte(transactionReport)
	content := r.mapCsv(transactionReportFile, bytes.NewBuffer(body))

	user := sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@example.com", time.Now(), time.Now())
	transactionTypeRows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, transactionType, time.Now(), time.Now())
	marketplaceRows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, marketplace, time.Now(), time.Now())
	fulfillmentRows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, fulfillment, time.Now(), time.Now())
	taxCollectionModelRows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, taxCollectionModel, time.Now(), time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	mock.ExpectQuery("^SELECT (.+) FROM transaction_types").WillReturnRows(transactionTypeRows)
	mock.ExpectQuery("^SELECT (.+) FROM marketplaces").WillReturnRows(marketplaceRows)
	mock.ExpectQuery("^SELECT (.+) FROM fulfillments").WillReturnRows(fulfillmentRows)
	mock.ExpectQuery("^SELECT (.+) FROM tax_collection_models").WillReturnRows(taxCollectionModelRows)
	mock.ExpectExec("^INSERT INTO transactions").WillReturnResult(sqlmock.NewResult(1, 1))

	formatted := r.formatTransactions(content, userId)
	err := r.saveTransaction(formatted[0])

	if err != nil {
		t.Error(err)
	}
}
