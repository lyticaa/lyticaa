package helpers

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
)

func BuildFilter(r *http.Request) *models.Filter {
	dateRange, startDate, endDate := DateRange(r)

	filter := models.NewFilter()
	filter.DateRange = dateRange
	filter.StartDate = startDate
	filter.EndDate = endDate
	filter.Start = DtStart(r)
	filter.Length = DtLength(r)
	filter.Dir = DtDir(r)
	filter.Sort = DtSort(r)

	return filter
}
