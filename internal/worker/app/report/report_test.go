package report

import (
	"bytes"
	"gitlab.com/getlytica/lytica-app/internal/models"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
)

const (
	typeCSV                    = "text/csv"
	typeXLSX                   = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	userId                     = "5de89aea5a61280de1f1bf2b"
	transactionReportFile      = "../../../../../lytica-app/test/fixtures/internal/worker/app/report/custom_transaction.csv"
	sponsoredProductReportFile = "../../../../../lytica-app/test/fixtures/internal/worker/app/report/sponsored_products.xlsx"
	currency                   = "USD"
	currencySymbol             = "$"
	transactionType            = "Order"
	marketplace                = "amazon.com"
	fulfillment                = "Amazon"
	taxCollectionModel         = "MarketplaceFacilitator"
)

type reportSuite struct {
	r    *Report
	mock sqlmock.Sqlmock
}

var _ = Suite(&reportSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *reportSuite) SetUpSuite(c *C) {
	dbMock, mock, err := sqlmock.New()
	c.Assert(err, IsNil)

	db := sqlx.NewDb(dbMock, "sqlmock")
	s.r = NewReport(db)

	s.mock = mock
}

func (s *reportSuite) TestFile(c *C) {
	content := s.r.toMap(typeCSV, s.readFile(transactionReportFile, c))
	c.Assert(assert.Greater(c, len(content), 0), Equals, true)

	content = s.r.toMap(typeXLSX, s.readFile(sponsoredProductReportFile, c))
	c.Assert(assert.Greater(c, len(content), 0), Equals, true)

	content = s.r.mapCSV(bytes.NewBuffer(s.readFile(transactionReportFile, c)))
	c.Assert(assert.Greater(c, len(content), 0), Equals, true)
	c.Assert(content[0]["date/time"], Equals, "Dec 1, 2019 12:07:47 AM PST")
	c.Assert(content[0]["settlement id"], Equals, "12447169531")
	c.Assert(content[0]["type"], Equals, "Order")
	c.Assert(content[0]["order id"], Equals, "113-0688349-7048213")
	c.Assert(content[0]["sku"], Equals, "PF-EV1C-1R5B")
	c.Assert(content[0]["description"], Equals, "Trained Flag Football Set,10 Man Set,Premium Football Gear, Massive 46 Piece Set, Flags, Belts, Cones, More, Bonus: Stylish Carry Bag & Flag Football")
	c.Assert(content[0]["quantity"], Equals, "1")
	c.Assert(content[0]["marketplace"], Equals, "amazon.com")
	c.Assert(content[0]["fulfillment"], Equals, "Amazon")
	c.Assert(content[0]["order city"], Equals, "Milford")
	c.Assert(content[0]["order state"], Equals, "DE")
	c.Assert(content[0]["tax collection model"], Equals, "")
	c.Assert(content[0]["product sales"], Equals, "26.5")
	c.Assert(content[0]["product sales tax"], Equals, "0")
	c.Assert(content[0]["shipping credits"], Equals, "0")
	c.Assert(content[0]["shipping credits tax"], Equals, "0")
	c.Assert(content[0]["gift wrap credits"], Equals, "0")
	c.Assert(content[0]["giftwrap credits tax"], Equals, "0")
	c.Assert(content[0]["promotional rebates"], Equals, "-0.27")
	c.Assert(content[0]["promotional rebates tax"], Equals, "0")
	c.Assert(content[0]["marketplace withheld tax"], Equals, "0")
	c.Assert(content[0]["selling fees"], Equals, "-3.93")
	c.Assert(content[0]["fba fees"], Equals, "-5.26")
	c.Assert(content[0]["other transaction fees"], Equals, "0")
	c.Assert(content[0]["other"], Equals, "0")
	c.Assert(content[0]["total"], Equals, "17.04")

	content = s.r.mapXLSX(s.readFile(sponsoredProductReportFile, c))
	c.Assert(assert.Greater(c, len(content), 0), Equals, true)
	c.Assert(content[0]["Start Date"], Equals, "Dec 01, 2019")
	c.Assert(content[0]["End Date"], Equals, "Dec 21, 2019")
	c.Assert(content[0]["Portfolio name"], Equals, "Not grouped")
	c.Assert(content[0]["Currency"], Equals, "USD")
	c.Assert(content[0]["Campaign Name"], Equals, "Flag Football Auto")
	c.Assert(content[0]["Ad Group Name"], Equals, "Ad Group 1")
	c.Assert(content[0]["Advertised SKU"], Equals, "PF-EV1C-1R5B")
	c.Assert(content[0]["Advertised ASIN"], Equals, "B01AQKSLMC")
	c.Assert(content[0]["Impressions"], Equals, "50293")
	c.Assert(content[0]["Clicks"], Equals, "47")
	c.Assert(content[0]["Click-Thru Rate (CTR)"], Equals, "0.0935%")
	c.Assert(content[0]["Cost Per Click (CPC)"], Equals, "$ 0.35")
	c.Assert(content[0]["Spend"], Equals, "$ 16.22")
	c.Assert(content[0]["7 Day Total Sales"], Equals, "$ 86.48")
	c.Assert(content[0]["Total Advertising Cost of Sales (ACoS)"], Equals, "18.7558%")
	c.Assert(content[0]["Total Return on Advertising Spend (RoAS)"], Equals, "5.33")
	c.Assert(content[0]["7 Day Total Orders (#)"], Equals, "3")
	c.Assert(content[0]["7 Day Total Units (#)"], Equals, "3")
	c.Assert(content[0]["7 Day Conversion Rate"], Equals, "6.3830%")
	c.Assert(content[0]["7 Day Advertised SKU Units (#)"], Equals, "3")
	c.Assert(content[0]["7 Day Other SKU Units (#)"], Equals, "1")
	c.Assert(content[0]["7 Day Advertised SKU Sales"], Equals, "$ 86.48")
	c.Assert(content[0]["7 Day Other SKU Sales"], Equals, "$ 1.00")
}

func (s *reportSuite) TestReport(c *C) {
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

	s.mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	s.mock.ExpectQuery("^SELECT (.+) FROM transaction_types").WillReturnRows(transactionTypeRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM marketplaces").WillReturnRows(marketplaceRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM fulfillments").WillReturnRows(fulfillmentRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM tax_collection_models").WillReturnRows(taxCollectionModelRows)
	s.mock.ExpectExec("^INSERT INTO transactions").WillReturnResult(sqlmock.NewResult(1, 1))

	body := s.readFile(transactionReportFile, c)
	err := s.r.processReport(customTransaction, userId, typeCSV, body)
	c.Assert(err, IsNil)

	exchangeRateRows := sqlmock.NewRows([]string{"id", "marketplace_id", "code", "symbol", "rate", "created_at", "updated_at"}).
		AddRow(1, 1, currency, currencySymbol, 1.00, time.Now(), time.Now())

	s.mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	s.mock.ExpectQuery("^SELECT (.+) FROM exchange_rates").WillReturnRows(exchangeRateRows)
	s.mock.ExpectExec("^INSERT INTO sponsored_products").WillReturnResult(sqlmock.NewResult(1, 1))

	body = s.readFile(sponsoredProductReportFile, c)
	err = s.r.processReport(sponsoredProducts, userId, typeXLSX, body)
	c.Assert(err, IsNil)

	user = sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@getlytica.com", time.Now(), time.Now())
	transactionTypeRows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, transactionType, time.Now(), time.Now())
	marketplaceRows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, marketplace, time.Now(), time.Now())
	fulfillmentRows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, fulfillment, time.Now(), time.Now())
	taxCollectionModelRows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, taxCollectionModel, time.Now(), time.Now())

	s.mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	s.mock.ExpectQuery("^SELECT (.+) FROM transaction_types").WillReturnRows(transactionTypeRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM marketplaces").WillReturnRows(marketplaceRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM fulfillments").WillReturnRows(fulfillmentRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM tax_collection_models").WillReturnRows(taxCollectionModelRows)
	s.mock.ExpectExec("^INSERT INTO transactions").WillReturnResult(sqlmock.NewResult(1, 1))

	content := s.r.mapCSV(bytes.NewBuffer(s.readFile(transactionReportFile, c)))
	c.Assert(assert.Greater(c, len(content), 0), Equals, true)

	err = s.r.processTransactions(content, userId)
	c.Assert(err, IsNil)

	user = sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@getlytica.com", time.Now(), time.Now())
	exchangeRateRows = sqlmock.NewRows([]string{"id", "marketplace_id", "code", "symbol", "rate", "created_at", "updated_at"}).
		AddRow(1, 1, currency, currencySymbol, 1.00, time.Now(), time.Now())

	s.mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	s.mock.ExpectQuery("^SELECT (.+) FROM exchange_rates").WillReturnRows(exchangeRateRows)
	s.mock.ExpectExec("^INSERT INTO sponsored_products").WillReturnResult(sqlmock.NewResult(1, 1))

	content = s.r.mapXLSX(s.readFile(sponsoredProductReportFile, c))
	c.Assert(assert.Greater(c, len(content), 0), Equals, true)

	err = s.r.processSponsoredProducts(content, userId)
	c.Assert(err, IsNil)

	actual := s.r.userFromFilename(userId + "/" + transactionReportFile)
	c.Assert(userId, Equals, actual)
}

func (s *reportSuite) TestSponsoredProduct(c *C) {
	rows := sqlmock.NewRows([]string{"id", "marketplace_id", "code", "symbol", "rate", "created_at", "updated_at"}).
		AddRow(1, 1, currency, currencySymbol, 1.0, time.Now(), time.Now())
	s.mock.ExpectQuery("^SELECT (.+) FROM exchange_rates").WillReturnRows(rows)

	exchangeRates := models.LoadExchangeRates(s.r.Db)
	c.Assert(len(exchangeRates), Equals, 1)
	c.Assert(exchangeRates[0].Code, Equals, currency)
	c.Assert(exchangeRates[0].Symbol, Equals, currencySymbol)

	rows = sqlmock.NewRows([]string{"id", "marketplace_id", "code", "symbol", "rate", "created_at", "updated_at"}).
		AddRow(1, 1, currency, currencySymbol, 1.0, time.Now(), time.Now())
	s.mock.ExpectQuery("^SELECT (.+) FROM exchange_rates").WillReturnRows(rows)

	exchangeRates = models.LoadExchangeRates(s.r.Db)
	_, ok := s.r.exchangeRate(currency, exchangeRates)
	c.Assert(ok, Equals, true)

	content := s.r.mapXLSX(s.readFile(sponsoredProductReportFile, c))

	user := sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@getlytica.com", time.Now(), time.Now())
	exchangeRateRows := sqlmock.NewRows([]string{"id", "marketplace_id", "code", "symbol", "rate", "created_at", "updated_at"}).
		AddRow(1, 1, currency, currencySymbol, 1.0, time.Now(), time.Now())

	s.mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	s.mock.ExpectQuery("^SELECT (.+) FROM exchange_rates").WillReturnRows(exchangeRateRows)

	formatted := s.r.formatSponsoredProducts(content, userId)
	c.Assert(len(formatted), Equals, 1)
	c.Assert(formatted[0].User.Id, Equals, int64(1))
	c.Assert(formatted[0].StartDate.IsZero(), Equals, false)
	c.Assert(formatted[0].EndDate.IsZero(), Equals, false)
	c.Assert(formatted[0].PortfolioName, Equals, "Not grouped")
	c.Assert(formatted[0].ExchangeRate.Id, Equals, int64(1))
	c.Assert(formatted[0].CampaignName, Equals, "Flag Football Auto")
	c.Assert(formatted[0].AdGroupName, Equals, "Ad Group 1")
	c.Assert(formatted[0].SKU, Equals, "PF-EV1C-1R5B")
	c.Assert(formatted[0].ASIN, Equals, "B01AQKSLMC")
	c.Assert(formatted[0].Impressions, Equals, int64(50293))
	c.Assert(formatted[0].Clicks, Equals, int64(47))
	c.Assert(formatted[0].CTR, Equals, 0.0935)
	c.Assert(formatted[0].CPC, Equals, 0.35)
	c.Assert(formatted[0].Spend, Equals, 16.22)
	c.Assert(formatted[0].TotalSales, Equals, 86.48)
	c.Assert(formatted[0].ACoS, Equals, 18.7558)
	c.Assert(formatted[0].RoAS, Equals, 5.33)
	c.Assert(formatted[0].TotalOrders, Equals, int64(3))
	c.Assert(formatted[0].TotalUnits, Equals, int64(3))
	c.Assert(formatted[0].ConversionRate, Equals, 6.383)
	c.Assert(formatted[0].AdvertisedSKUUnits, Equals, int64(3))
	c.Assert(formatted[0].OtherSKUUnits, Equals, int64(1))
	c.Assert(formatted[0].AdvertisedSKUSales, Equals, 86.48)
	c.Assert(formatted[0].OtherSKUSales, Equals, 1.0)

	content = s.r.mapXLSX(s.readFile(sponsoredProductReportFile, c))

	user = sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@getlytica.com", time.Now(), time.Now())
	exchangeRateRows = sqlmock.NewRows([]string{"id", "marketplace_id", "code", "symbol", "rate", "created_at", "updated_at"}).
		AddRow(1, 1, currency, currencySymbol, 1.0, time.Now(), time.Now())

	s.mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	s.mock.ExpectQuery("^SELECT (.+) FROM exchange_rates").WillReturnRows(exchangeRateRows)
	s.mock.ExpectExec("^INSERT INTO sponsored_products").WillReturnResult(sqlmock.NewResult(1, 1))

	formatted = s.r.formatSponsoredProducts(content, userId)
	err := s.r.saveSponsoredProduct(formatted[0])
	c.Assert(err, IsNil)
}

func (s *reportSuite) TestStorage(c *C) {}

func (s *reportSuite) TestTransaction(c *C) {
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, transactionType, time.Now(), time.Now())
	s.mock.ExpectQuery("^SELECT (.+) FROM transaction_types").WillReturnRows(rows)

	transactionTypes := s.r.getTransactionTypes()
	c.Assert(len(transactionTypes), Equals, 1)
	c.Assert(transactionTypes[0].Name, Equals, transactionType)

	rows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, transactionType, time.Now(), time.Now())
	s.mock.ExpectQuery("^SELECT (.+) FROM transaction_types").WillReturnRows(rows)

	transactionTypes = s.r.getTransactionTypes()
	_, ok := s.r.getTransactionTypeIdByName(transactionType, transactionTypes)
	c.Assert(ok, Equals, true)

	rows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, marketplace, time.Now(), time.Now())
	s.mock.ExpectQuery("^SELECT (.+) FROM marketplaces").WillReturnRows(rows)

	marketplaces := s.r.getMarketplaces()
	c.Assert(len(marketplaces), Equals, 1)
	c.Assert(marketplaces[0].Name, Equals, marketplace)

	rows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, marketplace, time.Now(), time.Now())
	s.mock.ExpectQuery("^SELECT (.+) FROM marketplaces").WillReturnRows(rows)

	marketplaces = s.r.getMarketplaces()
	_, ok = s.r.getMarketplaceIdByName(marketplace, marketplaces)
	c.Assert(ok, Equals, true)

	rows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, fulfillment, time.Now(), time.Now())
	s.mock.ExpectQuery("^SELECT (.+) FROM fulfillments").WillReturnRows(rows)

	fulfillments := s.r.getFulfillments()
	c.Assert(len(fulfillments), Equals, 1)
	c.Assert(marketplaces[0].Name, Equals, marketplace)

	rows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, fulfillment, time.Now(), time.Now())
	s.mock.ExpectQuery("^SELECT (.+) FROM fulfillments").WillReturnRows(rows)

	fulfillments = s.r.getFulfillments()
	_, ok = s.r.getFulfillmentIdByName(fulfillment, fulfillments)
	c.Assert(ok, Equals, true)

	rows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, taxCollectionModel, time.Now(), time.Now())
	s.mock.ExpectQuery("^SELECT (.+) FROM tax_collection_models").WillReturnRows(rows)

	taxCollectionModels := s.r.getTaxCollectionModels()
	c.Assert(len(taxCollectionModels), Equals, 1)
	c.Assert(taxCollectionModels[0].Name, Equals, taxCollectionModel)

	rows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, taxCollectionModel, time.Now(), time.Now())
	s.mock.ExpectQuery("^SELECT (.+) FROM tax_collection_models").WillReturnRows(rows)

	taxCollectionModels = s.r.getTaxCollectionModels()
	_, ok = s.r.getTaxCollectionModelIdByName(taxCollectionModel, taxCollectionModels)
	c.Assert(ok, Equals, true)

	content := s.r.mapCSV(bytes.NewBuffer(s.readFile(transactionReportFile, c)))

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

	s.mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	s.mock.ExpectQuery("^SELECT (.+) FROM transaction_types").WillReturnRows(transactionTypeRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM marketplaces").WillReturnRows(marketplaceRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM fulfillments").WillReturnRows(fulfillmentRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM tax_collection_models").WillReturnRows(taxCollectionModelRows)

	formatted := s.r.formatTransactions(content, userId)
	c.Assert(len(formatted), Equals, 1)
	c.Assert(formatted[0].User.Id, Equals, int64(1))
	c.Assert(formatted[0].DateTime.IsZero(), Equals, false)
	c.Assert(formatted[0].SettlementIdx, Equals, int64(1))
	c.Assert(formatted[0].SettlementId, Equals, int64(12447169531))
	c.Assert(formatted[0].TransactionType.Id, Equals, int64(1))
	c.Assert(formatted[0].OrderId, Equals, "113-0688349-7048213")
	c.Assert(formatted[0].SKU, Equals, "PF-EV1C-1R5B")
	c.Assert(formatted[0].Description, Equals, "Trained Flag Football Set,10 Man Set,Premium Football Gear, Massive 46 Piece Set, Flags, Belts, Cones, More, Bonus: Stylish Carry Bag & Flag Football")
	c.Assert(formatted[0].Quantity, Equals, int64(1))
	c.Assert(formatted[0].Marketplace.Id, Equals, int64(1))
	c.Assert(formatted[0].Fulfillment.Id, Equals, int64(1))
	c.Assert(formatted[0].TaxCollectionModel.Id, Equals, int64(1))
	c.Assert(formatted[0].ProductSales, Equals, 26.5)
	c.Assert(formatted[0].ProductSalesTax, Equals, 0.0)
	c.Assert(formatted[0].ShippingCredits, Equals, 0.0)
	c.Assert(formatted[0].ShippingCreditsTax, Equals, 0.0)
	c.Assert(formatted[0].GiftwrapCredits, Equals, 0.0)
	c.Assert(formatted[0].GiftwrapCreditsTax, Equals, 0.0)
	c.Assert(formatted[0].PromotionalRebates, Equals, -0.27)
	c.Assert(formatted[0].PromotionalRebatesTax, Equals, 0.0)
	c.Assert(formatted[0].MarketplaceWithheldTax, Equals, 0.0)
	c.Assert(formatted[0].SellingFees, Equals, -3.93)
	c.Assert(formatted[0].FBAFees, Equals, -5.26)
	c.Assert(formatted[0].OtherTransactionFees, Equals, 0.0)
	c.Assert(formatted[0].Other, Equals, 0.0)
	c.Assert(formatted[0].Total, Equals, 17.04)

	content = s.r.mapCSV(bytes.NewBuffer(s.readFile(transactionReportFile, c)))

	user = sqlmock.NewRows([]string{"id", "user_id", "email", "created_at", "updated_at"}).
		AddRow(1, userId, "test@getlytica.com", time.Now(), time.Now())
	transactionTypeRows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, transactionType, time.Now(), time.Now())
	marketplaceRows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, marketplace, time.Now(), time.Now())
	fulfillmentRows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, fulfillment, time.Now(), time.Now())
	taxCollectionModelRows = sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow(1, taxCollectionModel, time.Now(), time.Now())

	s.mock.ExpectQuery("^SELECT (.+) FROM users WHERE").WillReturnRows(user)
	s.mock.ExpectQuery("^SELECT (.+) FROM transaction_types").WillReturnRows(transactionTypeRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM marketplaces").WillReturnRows(marketplaceRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM fulfillments").WillReturnRows(fulfillmentRows)
	s.mock.ExpectQuery("^SELECT (.+) FROM tax_collection_models").WillReturnRows(taxCollectionModelRows)
	s.mock.ExpectExec("^INSERT INTO transactions").WillReturnResult(sqlmock.NewResult(1, 1))

	formatted = s.r.formatTransactions(content, userId)
	err := s.r.saveTransaction(formatted[0])
	c.Assert(err, IsNil)
}

func (s *reportSuite) TestTranslationHeader(c *C) {
	expected := "date/time"
	actual := s.r.translateHeader(expected)
	c.Assert(actual, Equals, expected)
}

func (s *reportSuite) TearDownSuite(c *C) {}

func (s *reportSuite) readFile(filename string, c *C) []byte {
	absPath, _ := filepath.Abs(filename)
	body, err := ioutil.ReadFile(absPath)
	c.Assert(err, IsNil)

	return body
}
