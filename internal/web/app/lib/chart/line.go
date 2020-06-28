package chart

import (
	"fmt"
	"strings"

	"gitlab.com/getlytica/lytica-app/internal/web/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

func (c *Chart) Line(data *[]types.Summary, dateRange string) types.Chart {
	var chart types.Chart

	c.categories(data, &chart, dateRange)
	c.series(data, &chart)

	return chart
}

func (c *Chart) categories(data *[]types.Summary, chart *types.Chart, dateRange string) {
	var categories []string
	for _, item := range *data {
		categories = append(categories, fmt.Sprintf("%v", helpers.DateFormat(dateRange, item.Date)))
	}

	chart.Line.Categories = append(chart.Line.Categories, types.Category{Category: strings.Join(categories, "|")})
}

func (c *Chart) series(data *[]types.Summary, chart *types.Chart) {
	series := make(map[string][]string)
	for _, item := range *data {
		series[item.Marketplace] = append(series[item.Marketplace], fmt.Sprintf("%v", item.Total))
	}

	for marketplace, item := range series {
		dataSet := types.DataSet{
			SeriesName: marketplace,
			Data:       strings.Join(item, "|"),
		}

		chart.Line.DataSets = append(chart.Line.DataSets, dataSet)
	}

	if len(chart.Line.DataSets) == 0 {
		chart.Line.DataSets = append(chart.Line.DataSets, types.DataSet{})
	}
}
