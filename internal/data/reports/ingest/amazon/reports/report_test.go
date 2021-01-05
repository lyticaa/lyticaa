package reports

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	transactions "github.com/lyticaa/lyticaa-app/internal/data/reports/ingest/amazon/reports/custom_transactions"
	sponsored "github.com/lyticaa/lyticaa-app/internal/data/reports/ingest/amazon/reports/sponsored_products"
	"github.com/lyticaa/lyticaa-app/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	. "gopkg.in/check.v1"
	"syreclabs.com/go/faker"
)

const (
	typeCSV                    = "text/csv"
	typeXLSX                   = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	transactionReportFile      = "../../../../../../test/fixtures/custom_transaction.csv"
	sponsoredProductReportFile = "../../../../../../test/fixtures/sponsored_products.xlsx"
	transactionType            = "Order"
	marketplace                = "amazon.com"
	fulfillment                = "Amazon"
	taxCollectionModel         = "MarketplaceFacilitator"
)

type reportsSuite struct {
	Reports     *Reports
	Sponsored   *sponsored.SponsoredProducts
	Transaction *transactions.CustomTransactions
}

var _ = Suite(&reportsSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *reportsSuite) SetUpSuite(c *C) {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	c.Assert(err, IsNil)

	s.Reports = NewReports(db)
	s.Sponsored = sponsored.NewSponsoredProducts(db)
	s.Transaction = transactions.NewCustomTransactions(db)
}

func (s *reportsSuite) TestFile(c *C) {
	amazonMarketplaceModel := models.AmazonMarketplaceModel{
		Name: marketplace,
	}

	amazonMarketplace := amazonMarketplaceModel.FetchOne(context.TODO(), s.Reports.db).(models.AmazonMarketplaceModel)
	c.Assert(amazonMarketplace.Name, Equals, marketplace)

	amazonFulfillmentModel := models.AmazonFulfillmentModel{
		Name: fulfillment,
	}

	amazonFulfilment := amazonFulfillmentModel.FetchOne(context.TODO(), s.Reports.db).(models.AmazonFulfillmentModel)
	c.Assert(amazonFulfilment.Name, Equals, fulfillment)

	amazonTaxCollectionModelModel := models.AmazonTaxCollectionModelModel{
		Name: taxCollectionModel,
	}

	amazonTaxCollectionModel := amazonTaxCollectionModelModel.FetchOne(context.TODO(), s.Reports.db).(models.AmazonTaxCollectionModelModel)
	c.Assert(amazonTaxCollectionModel.Name, Equals, taxCollectionModel)

	content := s.Reports.ToMap(typeCSV, s.readFile(transactionReportFile, c))
	c.Assert(assert.Greater(c, len(content), 0), Equals, true)

	content = s.Reports.ToMap(typeXLSX, s.readFile(sponsoredProductReportFile, c))
	c.Assert(assert.Greater(c, len(content), 0), Equals, true)

	content = s.Reports.MapCSV(bytes.NewBuffer(s.readFile(transactionReportFile, c)))
	c.Assert(assert.Greater(c, len(content), 0), Equals, true)
	c.Assert(content[0]["date/time"], Equals, "Dec 1, 2019 12:07:47 AM PST")
	c.Assert(content[0]["settlement id"], Equals, "12447169531")
	c.Assert(content[0]["type"], Equals, "Order")
	c.Assert(content[0]["order id"], Equals, "113-0688349-7048213")
	c.Assert(content[0]["sku"], Equals, "PF-EV1C-1R5B")
	c.Assert(content[0]["quantity"], Equals, "1")
	c.Assert(content[0]["marketplace"], Equals, amazonMarketplace.Name)
	c.Assert(content[0]["fulfillment"], Equals, amazonFulfilment.Name)
	c.Assert(content[0]["order city"], Equals, "Milford")
	c.Assert(content[0]["order state"], Equals, "DE")
	c.Assert(content[0]["tax collection model"], Equals, amazonTaxCollectionModel.Name)
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

	content = s.Reports.MapXLSX(s.readFile(sponsoredProductReportFile, c))
	c.Assert(assert.Greater(c, len(content), 0), Equals, true)
	c.Assert(content[0]["Date"], Equals, "Dec 01, 2019")
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

func (s *reportsSuite) TestSponsored(c *C) {
	content := s.Reports.MapXLSX(s.readFile(sponsoredProductReportFile, c))

	amazonMarketplaceModel := models.AmazonMarketplaceModel{
		Name: marketplace,
	}

	amazonMarketplace := amazonMarketplaceModel.FetchOne(context.TODO(), s.Reports.db).(models.AmazonMarketplaceModel)
	c.Assert(amazonMarketplace.Name, Equals, marketplace)

	userID := faker.RandomString(32)

	formatted := s.Sponsored.Format(content, userID)
	c.Assert(len(formatted), Equals, 1)
	c.Assert(formatted[0].UserID, Equals, userID)
	c.Assert(formatted[0].DateTime.IsZero(), Equals, false)
	c.Assert(formatted[0].PortfolioName, Equals, "Not grouped")
	c.Assert(formatted[0].AmazonMarketplaceID, Equals, amazonMarketplace.ID)
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

	err := s.Sponsored.Create(formatted[0])
	c.Assert(err, IsNil)
}

func (s *reportsSuite) TestTransaction(c *C) {
	amazonTransactionTypeModel := models.AmazonTransactionTypeModel{
		Name: transactionType,
	}

	amazonTransactionType := amazonTransactionTypeModel.FetchOne(context.TODO(), s.Reports.db).(models.AmazonTransactionTypeModel)
	c.Assert(amazonTransactionType.Name, Equals, transactionType)

	amazonMarketplaceModel := models.AmazonMarketplaceModel{
		Name: marketplace,
	}

	amazonMarketplace := amazonMarketplaceModel.FetchOne(context.TODO(), s.Reports.db).(models.AmazonMarketplaceModel)
	c.Assert(amazonMarketplace.Name, Equals, marketplace)

	amazonFulfillmentModel := models.AmazonFulfillmentModel{
		Name: fulfillment,
	}

	amazonFulfilment := amazonFulfillmentModel.FetchOne(context.TODO(), s.Reports.db).(models.AmazonFulfillmentModel)
	c.Assert(amazonFulfilment.Name, Equals, fulfillment)

	amazonTaxCollectionModelModel := models.AmazonTaxCollectionModelModel{
		Name: taxCollectionModel,
	}

	amazonTaxCollectionModel := amazonTaxCollectionModelModel.FetchOne(context.TODO(), s.Reports.db).(models.AmazonTaxCollectionModelModel)
	c.Assert(amazonTaxCollectionModel.Name, Equals, taxCollectionModel)

	content := s.Reports.MapCSV(bytes.NewBuffer(s.readFile(transactionReportFile, c)))

	userID := faker.RandomString(32)

	formatted := s.Transaction.Format(content, userID)
	c.Assert(len(formatted), Equals, 1)
	c.Assert(formatted[0].UserID, Equals, userID)
	c.Assert(formatted[0].DateTime.IsZero(), Equals, false)
	c.Assert(formatted[0].SettlementID, Equals, int64(12447169531))
	c.Assert(formatted[0].AmazonTransactionTypeID, Equals, amazonTransactionType.ID)
	c.Assert(formatted[0].OrderID, Equals, "113-0688349-7048213")
	c.Assert(formatted[0].SKU, Equals, "PF-EV1C-1R5B")
	c.Assert(formatted[0].Quantity, Equals, int64(1))
	c.Assert(formatted[0].AmazonMarketplaceID, Equals, amazonMarketplace.ID)
	c.Assert(formatted[0].AmazonFulfillmentID, Equals, amazonFulfilment.ID)
	c.Assert(formatted[0].AmazonTaxCollectionModelID, Equals, amazonTaxCollectionModel.ID)
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

	err := s.Transaction.Create(formatted[0])
	c.Assert(err, IsNil)
}

func (s *reportsSuite) TestReport(c *C) {
	userID := faker.RandomString(32)

	err := s.Reports.processReport(sponsoredProductReportFile, userID, typeXLSX, s.readFile(sponsoredProductReportFile, c))
	c.Assert(err, IsNil)

	err = s.Reports.processReport(transactionReportFile, userID, typeCSV, s.readFile(transactionReportFile, c))
	c.Assert(err, IsNil)
}

func (s *reportsSuite) TestStorage(c *C) {}

func (s *reportsSuite) TestTranslationHeader(c *C) {
	expected := "date/time"
	actual := s.Reports.translateHeader(expected)
	c.Assert(actual, Equals, expected)
}

func (s *reportsSuite) TearDownSuite(c *C) {}

func (s *reportsSuite) readFile(filename string, c *C) []byte {
	absPath, _ := filepath.Abs(filename)
	body, err := ioutil.ReadFile(absPath)
	c.Assert(err, IsNil)

	return body
}
