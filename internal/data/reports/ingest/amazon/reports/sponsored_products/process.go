package sponsored_products

import (
	"context"
	"github.com/lyticaa/lyticaa/internal/models"
)

const (
	unknown = int64(1)
)

func (s *SponsoredProducts) Process(rows []map[string]string, userID string) []error {
	var errors []error

	formatted := s.Format(rows, userID)
	for _, item := range formatted {
		if err := s.Create(item); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (s *SponsoredProducts) Format(rows []map[string]string, userID string) []models.AmazonSponsoredProductModel {
	var sponsoredProducts []models.AmazonSponsoredProductModel
	for _, row := range rows {
		sponsoredProducts = append(sponsoredProducts, models.AmazonSponsoredProductModel{
			UserID:              userID,
			DateTime:            s.dateTime(row),
			PortfolioName:       s.portfolioName(row),
			AmazonMarketplaceID: s.marketplaceID(row),
			CampaignName:        s.campaignName(row),
			AdGroupName:         s.adGroupName(row),
			SKU:                 s.sku(row),
			ASIN:                s.asin(row),
			Impressions:         s.impressions(row),
			Clicks:              s.clicks(row),
			CTR:                 s.ctr(row),
			CPC:                 s.cpc(row),
			Spend:               s.spend(row),
			TotalSales:          s.totalSales(row),
			ACoS:                s.acos(row),
			RoAS:                s.roas(row),
			TotalOrders:         s.totalOrders(row),
			TotalUnits:          s.totalUnits(row),
			ConversionRate:      s.conversionRate(row),
			AdvertisedSKUUnits:  s.advertisedSKUUnits(row),
			OtherSKUUnits:       s.otherSKUUnits(row),
			AdvertisedSKUSales:  s.advertisedSKUSales(row),
			OtherSKUSales:       s.otherSKUSales(row),
		})
	}

	return sponsoredProducts
}

func (s *SponsoredProducts) Create(sponsoredProduct models.AmazonSponsoredProductModel) error {
	if err := sponsoredProduct.Create(context.TODO(), s.db); err != nil {
		return err
	}

	return nil
}

func (s *SponsoredProducts) exchangeRates() *[]models.ExchangeRateModel {
	var exchangeRateModel models.ExchangeRateModel
	exchangeRates := exchangeRateModel.FetchAll(context.TODO(), nil, nil, s.db).([]models.ExchangeRateModel)

	return &exchangeRates
}
