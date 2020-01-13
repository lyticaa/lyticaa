package report

import (
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

	content := r.mapXlsx(readFile(sponsoredProductReportFile, t))

	user := sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@getlytica.com", time.Now(), time.Now())
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
			t.Error()
		}

		if formatted[0].StartDate.IsZero() {
			t.Error()
		}

		if formatted[0].EndDate.IsZero() {
			t.Error()
		}

		if formatted[0].PortfolioName != "Not grouped" {
			t.Error()
		}

		if formatted[0].Currency.Id != 1 {
			t.Error()
		}

		if formatted[0].CampaignName != "Flag Football Auto" {
			t.Error()
		}

		if formatted[0].AdGroupName != "Ad Group 1" {
			t.Error()
		}

		if formatted[0].SKU != "PF-EV1C-1R5B" {
			t.Error()
		}

		if formatted[0].ASIN != "B01AQKSLMC" {
			t.Error()
		}

		if formatted[0].Impressions != 50293 {
			t.Error()
		}

		if formatted[0].Clicks != 47 {
			t.Error()
		}

		if formatted[0].CTR != 0.0935 {
			t.Error()
		}

		if formatted[0].CPC != 0.35 {
			t.Error()
		}

		if formatted[0].Spend != 16.22 {
			t.Error()
		}

		if formatted[0].TotalSales != 86.48 {
			t.Error()
		}

		if formatted[0].ACoS != 18.7558 {
			t.Error()
		}

		if formatted[0].RoAS != 5.33 {
			t.Error()
		}

		if formatted[0].TotalOrders != 3 {
			t.Error()
		}

		if formatted[0].TotalUnits != 3 {
			t.Error()
		}

		if formatted[0].ConversionRate != 6.383 {
			t.Error()
		}

		if formatted[0].AdvertisedSKUUnits != 3 {
			t.Error()
		}

		if formatted[0].OtherSKUUnits != 1 {
			t.Error()
		}

		if formatted[0].AdvertisedSKUSales != 86.48 {
			t.Error()
		}

		if formatted[0].OtherSKUSales != 1.0 {
			t.Error()
		}
	}
}

func TestSaveSponsoredProduct(t *testing.T) {
	r, mock, complete := SetupTests(t)
	defer complete(r)

	content := r.mapXlsx(readFile(sponsoredProductReportFile, t))

	user := sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@getlytica.com", time.Now(), time.Now())
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
