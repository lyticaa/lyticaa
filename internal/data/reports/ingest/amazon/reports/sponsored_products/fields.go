package sponsored_products

import (
	"context"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/lyticaa/lyticaa/internal/models"
)

var (
	reg, _ = regexp.Compile(`[^0-9\.]+`)
)

func (s *SponsoredProducts) dateTime(row map[string]string) time.Time {
	var dateTime time.Time

	if days, err := strconv.Atoi(row["Date"]); err == nil {
		loc, _ := time.LoadLocation(os.Getenv("TZ"))
		dateTime = time.Date(1900, time.January, days, 0, 0, 0, 0, loc)
	} else {
		dateTime, _ = time.Parse("Jan 2, 2006", row["Date"])
	}

	return dateTime
}

func (s *SponsoredProducts) portfolioName(row map[string]string) string {
	return row["Portfolio name"]
}

func (s *SponsoredProducts) marketplaceID(row map[string]string) int64 {
	exchangeRateID := s.exchangeRateID(row)
	var amazonMarketplaceModel models.AmazonMarketplaceModel
	amazonMarketplaces := amazonMarketplaceModel.FetchAll(context.TODO(), nil, nil, s.db).([]models.AmazonMarketplaceModel)

	for _, marketplace := range amazonMarketplaces {
		if marketplace.ExchangeRateID == exchangeRateID {
			return marketplace.ID
		}
	}

	return unknown
}

func (s *SponsoredProducts) exchangeRateID(row map[string]string) int64 {
	for _, exchangeRate := range *s.exchangeRates() {
		if exchangeRate.Code == row["Currency"] {
			return exchangeRate.ID
		}
	}

	return unknown
}

func (s *SponsoredProducts) campaignName(row map[string]string) string {
	return row["Campaign Name"]
}

func (s *SponsoredProducts) adGroupName(row map[string]string) string {
	return row["Ad Group Name"]
}

func (s *SponsoredProducts) sku(row map[string]string) string {
	return row["Advertised SKU"]
}

func (s *SponsoredProducts) asin(row map[string]string) string {
	return row["Advertised ASIN"]
}

func (s *SponsoredProducts) impressions(row map[string]string) int64 {
	impressions, _ := strconv.ParseInt(row["Impressions"], 10, 64)
	return impressions
}

func (s *SponsoredProducts) clicks(row map[string]string) int64 {
	clicks, _ := strconv.ParseInt(row["Clicks"], 10, 64)
	return clicks
}

func (s *SponsoredProducts) ctr(row map[string]string) float64 {
	ctr, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Click-Thru Rate (CTR)"], ""), 64)
	return ctr
}

func (s *SponsoredProducts) cpc(row map[string]string) float64 {
	cpc, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Cost Per Click (CPC)"], ""), 64)
	return cpc
}

func (s *SponsoredProducts) spend(row map[string]string) float64 {
	spend, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Spend"], ""), 64)
	return spend
}

func (s *SponsoredProducts) totalSales(row map[string]string) float64 {
	totalSales, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Total Sales"], ""), 64)
	return totalSales
}

func (s *SponsoredProducts) acos(row map[string]string) float64 {
	acos, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Total Advertising Cost of Sales (ACoS)"], ""), 64)
	return acos
}

func (s *SponsoredProducts) roas(row map[string]string) float64 {
	roas, _ := strconv.ParseFloat(row["Total Return on Advertising Spend (RoAS)"], 64)
	return roas
}

func (s *SponsoredProducts) totalOrders(row map[string]string) int64 {
	totalOrders, _ := strconv.ParseInt(row["7 Day Total Orders (#)"], 10, 64)
	return totalOrders
}

func (s *SponsoredProducts) totalUnits(row map[string]string) int64 {
	totalUnits, _ := strconv.ParseInt(row["7 Day Total Units (#)"], 10, 64)
	return totalUnits
}

func (s *SponsoredProducts) conversionRate(row map[string]string) float64 {
	conversionRate, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Conversion Rate"], ""), 64)
	return conversionRate
}

func (s *SponsoredProducts) advertisedSKUUnits(row map[string]string) int64 {
	advertisedSKUUnits, _ := strconv.ParseInt(row["7 Day Advertised SKU Units (#)"], 10, 64)
	return advertisedSKUUnits
}

func (s *SponsoredProducts) otherSKUUnits(row map[string]string) int64 {
	otherSKUUnits, _ := strconv.ParseInt(row["7 Day Other SKU Units (#)"], 10, 64)
	return otherSKUUnits
}

func (s *SponsoredProducts) advertisedSKUSales(row map[string]string) float64 {
	advertisedSKUSales, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Advertised SKU Sales"], ""), 64)
	return advertisedSKUSales
}

func (s *SponsoredProducts) otherSKUSales(row map[string]string) float64 {
	otherSKUSales, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Other SKU Sales"], ""), 64)
	return otherSKUSales
}
