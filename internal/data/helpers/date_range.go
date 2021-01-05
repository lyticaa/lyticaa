package helpers

import (
	"time"
)

var (
	rangeToday               = "today"
	rangeYesterday           = "yesterday"
	rangeLastThirtyDays      = "last_thirty_days"
	rangePreviousThirtyDays  = "previous_thirty_days"
	rangeThisMonth           = "this_month"
	rangeLastMonth           = "last_month"
	rangeMonthBeforeLast     = "month_before_last"
	rangeLastThreeMonths     = "last_three_months"
	rangePreviousThreeMonths = "previous_three_months"
	rangeLastSixMonths       = "last_six_months"
	rangePreviousSixMonths   = "previous_six_months"
	rangeThisYear            = "this_year"
	rangeLastYear            = "last_year"
	rangeAllTime             = "all_time"
)

var (
	Ranges = []string{
		rangeToday,
		rangeYesterday,
		rangeLastThirtyDays,
		rangePreviousThirtyDays,
		rangeThisMonth,
		rangeLastMonth,
		rangeMonthBeforeLast,
		rangeLastThreeMonths,
		rangePreviousThreeMonths,
		rangeLastSixMonths,
		rangePreviousSixMonths,
		rangeThisYear,
		rangeLastYear,
		rangeAllTime,
	}
)

func DateFormat(dateRange string, date time.Time) string {
	var layout string
	switch dateRange {
	case rangeAllTime:
		layout = "Jan 2006"
	default:
		layout = "Jan 2, 2006 at 3:04pm (MST)"
	}

	return date.Format(layout)
}

func DateRange(dateRange string) (string, time.Time, time.Time) {
	var start time.Time
	var end time.Time

	switch dateRange {
	case rangeToday:
		start, end = today()
	case rangeLastThirtyDays:
		start, end = lastThirtyDays()
	case rangeThisMonth:
		start, end = thisMonth()
	case rangeLastMonth:
		start, end = lastMonth()
	case rangeLastThreeMonths:
		start, end = lastThreeMonths()
	case rangeLastSixMonths:
		start, end = lastSixMonths()
	case rangeThisYear:
		start, end = thisYear()
	case rangeAllTime:
		start, end = allTime()
	default:
		start, end = today()
	}

	return dateRange, start, end
}

func PreviousDateRange(dateRange string) (time.Time, time.Time) {
	switch dateRange {
	case rangeToday:
		return yesterday()
	case rangeLastThirtyDays:
		return previousThirtyDays()
	case rangeThisMonth:
		return lastMonth()
	case rangeLastMonth:
		return monthBeforeLast()
	case rangeLastThreeMonths:
		return previousThreeMonths()
	case rangeLastSixMonths:
		return previousSixMonths()
	case rangeThisYear:
		return lastYear()
	case rangeAllTime:
		return allTime()
	}

	return yesterday()
}

func PreviousDateRangeLabel(dateRange string) string {
	switch dateRange {
	case rangeToday:
		return rangeYesterday
	case rangeLastThirtyDays:
		return rangePreviousThirtyDays
	case rangeThisMonth:
		return rangeLastMonth
	case rangeLastMonth:
		return rangeMonthBeforeLast
	case rangeLastThreeMonths:
		return rangePreviousThreeMonths
	case rangeLastSixMonths:
		return rangePreviousSixMonths
	case rangeThisYear:
		return rangeLastYear
	}

	return dateRange
}

func IsDateRangeAllTime(dateRange string) bool {
	allTime := false

	if dateRange == rangeAllTime {
		allTime = true
	}

	return allTime
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

func monthBeforeLast() (time.Time, time.Time) {
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
