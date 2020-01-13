package report

import (
	"bytes"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

const (
	typeCsv                    = "text/csv"
	typeXlsx                   = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	userId                     = "5de89aea5a61280de1f1bf2b"
	transactionReportFile      = "../../../../../lytica/test/fixtures/internal/worker/app/report/custom_transaction.csv"
	sponsoredProductReportFile = "../../../../../lytica/test/fixtures/internal/worker/app/report/sponsored_products.xlsx"
)

func SetupTests(t *testing.T) (*Report, sqlmock.Sqlmock, func(*Report)) {
	dbM, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbM, "sqlmock")
	r := NewReport(db)

	return r, mock, func(r *Report) {
		_ = r.Db.Close()
	}
}

func TestProcessReport(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	user := sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@getlytica.com", time.Now(), time.Now())

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

	body := readFile(transactionReportFile, t)
	err := r.processReport(customTransaction, userId, typeCsv, body)
	if err != nil {
		t.Error(err)
	}

	currencyRows := sqlmock.NewRows([]string{"id", "name", "symbol", "created_at", "updated_at"}).
		AddRow(1, currency, currencySymbol, time.Now(), time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	mock.ExpectQuery("^SELECT (.+) FROM currencies").WillReturnRows(currencyRows)
	mock.ExpectExec("^INSERT INTO sponsored_products").WillReturnResult(sqlmock.NewResult(1, 1))

	body = readFile(sponsoredProductReportFile, t)
	err = r.processReport(sponsoredProducts, userId, typeXlsx, body)
	if err != nil {
		t.Error(err)
	}
}

func TestProcessTransactions(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	user := sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@getlytica.com", time.Now(), time.Now())
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

	content := r.mapCsv(bytes.NewBuffer(readFile(transactionReportFile, t)))
	if len(content) == 0 {
		t.Error("no rows generated")
	}

	err := r.processTransactions(content, userId)
	if err != nil {
		t.Error(err)
	}
}

func TestProcessSponsoredProducts(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	user := sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@getlytica.com", time.Now(), time.Now())
	currencyRows := sqlmock.NewRows([]string{"id", "name", "symbol", "created_at", "updated_at"}).
		AddRow(1, currency, currencySymbol, time.Now(), time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	mock.ExpectQuery("^SELECT (.+) FROM currencies").WillReturnRows(currencyRows)
	mock.ExpectExec("^INSERT INTO sponsored_products").WillReturnResult(sqlmock.NewResult(1, 1))

	content := r.mapXlsx(readFile(sponsoredProductReportFile, t))
	if len(content) == 0 {
		t.Error("no rows generated")
	}

	err := r.processSponsoredProducts(content, userId)
	if err != nil {
		t.Error(err)
	}
}

func TestUserFromFilename(t *testing.T) {
	r, _, complete := SetupTests(t)
	defer complete(r)

	actual := r.userFromFilename(userId + "/" + transactionReportFile)

	if actual != userId {
		t.Errorf("expected a user Id of %v but got %v instead!", userId, actual)
	}
}
