package data

import (
	"github.com/lyticaa/lyticaa-app/internal/web/types"
)

func (d *Data) Forecast(userID, dateRange, view string, forecast *types.Forecast) {
	forecast.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userID, dateRange, view), view), dateRange)
}
