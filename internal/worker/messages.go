package app

import (
	"encoding/json"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/worker/types"

	"github.com/streadway/amqp"
)

func (a *App) parse(msg amqp.Delivery) {
	var cd types.Payload
	if err := json.Unmarshal(msg.Body, &cd); err != nil {
		a.Logger.Error().Err(err).Msg("failed to unmarshal message body")
		return
	}

	a.Logger.Info().Msgf("processing message for %v", cd.Op)

	a.publish(cd.Body)
}

func (a *App) publish(data string) {
	var publishData types.Data
	if err := json.Unmarshal([]byte(data), &publishData); err != nil {
		a.Logger.Error().Err(err).Msg("failed to unmarshal message body")
		return
	}

	a.overview(publishData.Overview)
}

func (a *App) overview(data []types.Parsed) {
	for _, row := range data {
		dashboard := models.LoadDashboardByMarketplace(row.UserId, row.DateRange, row.Marketplace, row.DateTime, a.Db)
		dashboard.TotalSales = row.TotalSales
		dashboard.UnitsSold = row.UnitsSold
		dashboard.AmazonCosts = row.AmazonCosts
		dashboard.ProductCosts = row.ProductCosts
		dashboard.AdvertisingSpend = row.AdvertisingSpend
		dashboard.Refunds = row.Refunds
		dashboard.ShippingCredits = row.ShippingCredits
		dashboard.PromotionalRebates = row.PromotionalRebates
		dashboard.TotalCosts = row.TotalCosts
		dashboard.GrossMargin = row.GrossMargin
		dashboard.NetMargin = row.NetMargin
		_ = dashboard.Save(a.Db)
	}
}
