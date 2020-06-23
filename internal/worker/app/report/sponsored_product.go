package report

import (
	"regexp"
	"strconv"
	"time"

	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (r *Report) exchangeRate(code string, exchangeRates []models.ExchangeRate) (int64, bool) {
	for _, rate := range exchangeRates {
		if rate.Code == code {
			return rate.Id, true
		}
	}

	return unknown, false
}

func (r *Report) formatSponsoredProducts(rows []map[string]string, username string) []models.SponsoredProduct {
	user := models.LoadUser(username, r.Db)
	exchangeRates := models.LoadExchangeRates(r.Db)

	var sponsoredProducts []models.SponsoredProduct
	for _, row := range rows {
		// Remove all currency/percentage characters
		reg, _ := regexp.Compile(`[^0-9\.]+`)

		exchangeRate, ok := r.exchangeRate(row["Currency"], exchangeRates)
		if !ok && row["Currency"] != "" {
			r.Logger.Error().Msgf("Currency %v not found", row["Currency"])
		}

		startDate, _ := time.Parse("Jan 2, 2006", row["Start Date"])
		endDate, _ := time.Parse("Jan 2, 2006", row["End Date"])
		impressions, _ := strconv.ParseInt(row["Impressions"], 10, 64)
		clicks, _ := strconv.ParseInt(row["Clicks"], 10, 64)
		ctr, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Click-Thru Rate (CTR)"], ""), 64)
		cpc, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Cost Per Click (CPC)"], ""), 64)
		spend, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Spend"], ""), 64)
		totalSales, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Total Sales"], ""), 64)
		acos, _ := strconv.ParseFloat(reg.ReplaceAllString(row["Total Advertising Cost of Sales (ACoS)"], ""), 64)
		roas, _ := strconv.ParseFloat(row["Total Return on Advertising Spend (RoAS)"], 64)
		totalOrders, _ := strconv.ParseInt(row["7 Day Total Orders (#)"], 10, 64)
		totalUnits, _ := strconv.ParseInt(row["7 Day Total Units (#)"], 10, 64)
		conversionRate, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Conversion Rate"], ""), 64)
		advertisedSKUUnits, _ := strconv.ParseInt(row["7 Day Advertised SKU Units (#)"], 10, 64)
		otherSKUUnits, _ := strconv.ParseInt(row["7 Day Other SKU Units (#)"], 10, 64)
		advertisedSKUSales, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Advertised SKU Sales"], ""), 64)
		otherSKUSales, _ := strconv.ParseFloat(reg.ReplaceAllString(row["7 Day Other SKU Sales"], ""), 64)

		sponsoredProduct := models.SponsoredProduct{
			User:               *user,
			StartDate:          startDate,
			EndDate:            endDate,
			PortfolioName:      row["Portfolio name"],
			ExchangeRate:       models.ExchangeRate{Id: exchangeRate},
			CampaignName:       row["Campaign Name"],
			AdGroupName:        row["Ad Group Name"],
			SKU:                row["Advertised SKU"],
			ASIN:               row["Advertised ASIN"],
			Impressions:        impressions,
			Clicks:             clicks,
			CTR:                ctr,
			CPC:                cpc,
			Spend:              spend,
			TotalSales:         totalSales,
			ACoS:               acos,
			RoAS:               roas,
			TotalOrders:        totalOrders,
			TotalUnits:         totalUnits,
			ConversionRate:     conversionRate,
			AdvertisedSKUUnits: advertisedSKUUnits,
			OtherSKUUnits:      otherSKUUnits,
			AdvertisedSKUSales: advertisedSKUSales,
			OtherSKUSales:      otherSKUSales,
		}

		sponsoredProducts = append(sponsoredProducts, sponsoredProduct)
	}

	return sponsoredProducts
}

func (r *Report) saveSponsoredProduct(sponsoredProduct models.SponsoredProduct) error {
	err := models.SaveSponsoredProduct(sponsoredProduct, r.Db)
	if err != nil {
		return err
	}

	return nil
}
