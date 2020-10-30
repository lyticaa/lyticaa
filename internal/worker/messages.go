package worker

import (
	"encoding/json"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/worker/types"

	"github.com/streadway/amqp"
)

func (w *Worker) parse(msg amqp.Delivery) {
	var cd types.Payload
	if err := json.Unmarshal(msg.Body, &cd); err != nil {
		w.Logger.Error().Err(err).Msg("failed to unmarshal message body")
		return
	}

	w.Logger.Info().Msgf("processing message for %v", cd.Op)

	w.publish(cd.Body)
}

func (w *Worker) publish(data string) {
	var publishData types.Data
	if err := json.Unmarshal([]byte(data), &publishData); err != nil {
		w.Logger.Error().Err(err).Msg("failed to unmarshal message body")
		return
	}

	w.overview(publishData.Overview)
}

func (w *Worker) overview(data []types.Parsed) {
	for _, row := range data {
		dashboard := models.LoadDashboardByMarketplace(row.UserID, row.DateRange, row.Marketplace, row.DateTime, w.Db)
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
		_ = dashboard.Save(w.Db)
	}
}
