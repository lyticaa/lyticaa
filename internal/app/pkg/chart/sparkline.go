package chart

import (
	"github.com/lyticaa/lyticaa-app/internal/app/types"
)

func (c *Chart) Sparkline(data *[]types.Summary) types.Chart {
	var chart types.Chart
	for _, item := range *data {
		chart.Sparkline.Data = append(chart.Sparkline.Data, types.SparklineData{Value: item.Total})
	}

	if len(chart.Sparkline.Data) == 0 {
		chart.Sparkline.Data = []types.SparklineData{}
	}

	return chart
}
