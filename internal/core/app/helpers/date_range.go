package helpers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func DateRange(r *http.Request) (time.Time, time.Time) {
	params := mux.Vars(r)
	dateRange := params["dateRange"]
	switch dateRange {
	case "today":
		return today()
	case "last_30_days":
		return lastThirtyDays()
	case "this_month":
		return thisMonth()
	case "last_month":
		return lastMonth()
	case "last_3_months":
		return lastThreeMonths()
	case "last_6_months":
		return lastSixMonths()
	case "this_year":
		return thisYear()
	case "all_time":
		return allTime()
	}

	return today()
}

func PreviousDateRange(r *http.Request) (time.Time, time.Time) {
	params := mux.Vars(r)
	dateRange := params["dateRange"]
	switch dateRange {
	case "today":
		return yesterday()
	case "last_30_days":
		return previousThirtyDays()
	case "this_month":
		return lastMonth()
	case "last_month":
		return previousMonth()
	case "last_3_months":
		return previousThreeMonths()
	case "last_6_months":
		return previousSixMonths()
	case "this_year":
		return lastYear()
	case "all_time":
		return allTime()
	}

	return yesterday()
}

func today() (time.Time, time.Time) {
	year, month, day := time.Now().Date()
	return startDate(year, month, day), time.Now()
}

func yesterday() (time.Time, time.Time) {
	year, month, day := dateByDays(-1).Date()
	return startDate(year, month, day), endDate(year, month, day)
}

func lastThirtyDays() (time.Time, time.Time) {
	year, month, day := dateByDays(-30).Date()
	return startDate(year, month, day), time.Now()
}

func previousThirtyDays() (time.Time, time.Time) {
	year, month, day := dateByDays(-60).Date()
	return startDate(year, month, day), time.Now().AddDate(0, 0, -30)
}

func thisMonth() (time.Time, time.Time) {
	year, month, _ := time.Now().Date()
	return startDate(year, month, 1), time.Now()
}

func lastMonth() (time.Time, time.Time) {
	year, month, day := dateByMonths(-1, 0).Date()
	return startDate(year, month, day), time.Now()
}

func previousMonth() (time.Time, time.Time) {
	year, month, day := dateByMonths(-2, 0).Date()
	return startDate(year, month, day), time.Now().AddDate(0, -2, 0)
}

func lastThreeMonths() (time.Time, time.Time) {
	year, month, day := dateByMonths(-3, 0).Date()
	return startDate(year, month, day), time.Now()
}

func previousThreeMonths() (time.Time, time.Time) {
	year, month, day := dateByMonths(-6, 0).Date()
	return startDate(year, month, day), time.Now().AddDate(0, -3, 0)
}

func lastSixMonths() (time.Time, time.Time) {
	year, month, day := dateByMonths(-6, 0).Date()
	return startDate(year, month, day), time.Now()
}

func previousSixMonths() (time.Time, time.Time) {
	year, month, day := dateByMonths(-12, 0).Date()
	return startDate(year, month, day), time.Now().AddDate(0, -6, 0)
}

func thisYear() (time.Time, time.Time) {
	year, _, _ := time.Now().Date()
	sY, sM, sD := dateYear(year).Date()
	return startDate(sY, sM, sD), time.Now()
}

func lastYear() (time.Time, time.Time) {
	return time.Now().AddDate(-1, 0, 0), time.Now()
}

func allTime() (time.Time, time.Time) {
	return dateYear(1970), time.Now()
}

func startDate(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location())
}

func endDate(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 23, 59, 59, 0, time.Now().Location())
}

func dateByDays(day int) time.Time {
	return time.Now().AddDate(0, 0, day)
}

func dateByMonths(month int, day int) time.Time {
	return time.Now().AddDate(0, month, day)
}

func dateYear(year int) time.Time {
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.Now().Location())
}
