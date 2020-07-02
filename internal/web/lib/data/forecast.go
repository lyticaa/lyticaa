package data

import (
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (d *Data) Forecast(userId, dateRange, view string, forecast *types.Forecast) {
	forecast.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, view), view), dateRange)
}
