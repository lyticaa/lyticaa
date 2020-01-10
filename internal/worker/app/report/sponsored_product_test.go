package report

import (
	"bytes"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

const (
	currency       = "USD"
	currencySymbol = "$"
)

func TestGetCurrencies(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	rows := sqlmock.NewRows([]string{"id", "name", "symbol", "created_at", "updated_at"}).
		AddRow(1, currency, currencySymbol, time.Now(), time.Now())
	mock.ExpectQuery("^SELECT (.+) FROM currencies").WillReturnRows(rows)

	currencies := r.getCurrencies()
	if len(currencies) == 0 {
		t.Errorf("no currencies found!")
	}

	if len(currencies) > 0 {
		if currencies[0].Name != currency {
			t.Errorf("currency is invalid")
		}

		if currencies[0].Symbol != currencySymbol {
			t.Errorf("symbol is invalid")
		}
	}
}

func TestGetCurrencyIdByName(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	rows := sqlmock.NewRows([]string{"id", "name", "symbol", "created_at", "updated_at"}).
		AddRow(1, currency, currencySymbol, time.Now(), time.Now())
	mock.ExpectQuery("^SELECT (.+) FROM currencies").WillReturnRows(rows)

	currencies := r.getCurrencies()
	_, ok := r.getCurrencyIdByName(currency, currencies)
	if !ok {
		t.Errorf("currency not found")
	}
}

func TestFormatSponsoredProducts(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	body := []byte(sponsoredProductsReport)
	content := r.mapCsv(SponsoredProductReportFile, bytes.NewBuffer(body))

	user := sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@example.com", time.Now(), time.Now())
	currencyRows := sqlmock.NewRows([]string{"id", "name", "symbol", "created_at", "updated_at"}).
		AddRow(1, currency, currencySymbol, time.Now(), time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	mock.ExpectQuery("^SELECT (.+) FROM currencies").WillReturnRows(currencyRows)

	formatted := r.formatSponsoredProducts(content, userId)

	if len(formatted) == 0 {
		t.Errorf("no formatted sponsored products found!")
	}

	if len(formatted) > 0 {
		if formatted[0].User.Id != 1 {
			t.Errorf("invalid user Id")
		}

		if formatted[0].StartDate.IsZero() {
			t.Errorf("start date is invalid")
		}

		if formatted[0].EndDate.IsZero() {
			t.Errorf("end date is invalid")
		}

		if formatted[0].PortfolioName == "" {
			t.Errorf("portfolio name is invalid")
		}

		if formatted[0].Currency.Id != 1 {
			t.Errorf("currency is invalid")
		}

		if formatted[0].Impressions == 0 {
			t.Errorf("impressions are invalid")
		}

		if formatted[0].Clicks == 0 {
			t.Errorf("clicks are invalid")
		}

		if formatted[0].CTR == 0 {
			t.Errorf("click through rate is invalid")
		}

		if formatted[0].CPC == 0 {
			t.Errorf("cost per click is invalid")
		}

		if formatted[0].Spend == 0 {
			t.Errorf("spend is invalid")
		}
	}
}

func TestSaveSponsoredProduct(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	body := []byte(sponsoredProductsReport)
	content := r.mapCsv(SponsoredProductReportFile, bytes.NewBuffer(body))

	user := sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@example.com", time.Now(), time.Now())
	currencyRows := sqlmock.NewRows([]string{"id", "name", "symbol", "created_at", "updated_at"}).
		AddRow(1, currency, currencySymbol, time.Now(), time.Now())

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	mock.ExpectQuery("^SELECT (.+) FROM currencies").WillReturnRows(currencyRows)
	mock.ExpectExec("^INSERT INTO sponsored_products").WillReturnResult(sqlmock.NewResult(1, 1))

	formatted := r.formatSponsoredProducts(content, userId)
	err := r.saveSponsoredProduct(formatted[0])

	if err != nil {
		t.Error(err)
	}
}
