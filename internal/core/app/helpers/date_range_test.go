package helpers

import (
	"testing"
	"time"
)

const (
	epoch = 1970
)

func TestToday(t *testing.T) {
	start, end := today()

	date := time.Now()
	if start.Day() != date.Day() {
		t.Error()
	}

	if end.Day() != date.Day() {
		t.Error()
	}
}

func TestYesterday(t *testing.T) {
	start, end := yesterday()

	date := time.Now().AddDate(0, 0, -1)
	if start.Day() != date.Day() {
		t.Error()
	}

	if end.Day() != date.Day() {
		t.Error()
	}
}

func TestLastThirtyDays(t *testing.T) {
	start, end := lastThirtyDays()

	date := time.Now().AddDate(0, 0, -30)
	if start.Day() != date.Day() {
		t.Error()
	}

	if end.Day() != time.Now().Day() {
		t.Error()
	}
}

func TestPreviousThirtyDays(t *testing.T) {
	start, end := previousThirtyDays()

	date := time.Now().AddDate(0, 0, -60)
	if start.Day() != date.Day() {
		t.Error()
	}

	date = time.Now().AddDate(0, 0, -30)
	if end.Day() != date.Day() {
		t.Error()
	}
}

func TestThisMonth(t *testing.T) {
	start, end := thisMonth()

	date := time.Now()
	if start.Month() != date.Month() {
		t.Error()
	}

	if end.Month() != date.Month() {
		t.Error()
	}
}

func TestLastMonth(t *testing.T) {
	start, end := lastMonth()

	date := time.Now().AddDate(0, -1, 0)
	if start.Month() != date.Month() {
		t.Error()
	}

	if end.Month() != time.Now().Month() {
		t.Error()
	}
}

func TestPreviousMonth(t *testing.T) {
	start, end := previousMonth()

	date := time.Now().AddDate(0, -2, 0)
	if start.Month() != date.Month() {
		t.Error()
	}

	if end.Month() != date.Month() {
		t.Error()
	}
}

func TestLastThreeMonths(t *testing.T) {
	start, end := lastThreeMonths()

	date := time.Now().AddDate(0, -3, 0)
	if start.Month() != date.Month() {
		t.Error()
	}

	if end.Month() != time.Now().Month() {
		t.Error()
	}
}

func TestPreviousThreeMonths(t *testing.T) {
	start, end := previousThreeMonths()

	date := time.Now().AddDate(0, -6, 0)
	if start.Month() != date.Month() {
		t.Error()
	}

	date = time.Now().AddDate(0, -3, 0)
	if end.Month() != date.Month() {
		t.Error()
	}
}

func TestLastSixMonths(t *testing.T) {
	start, end := lastSixMonths()

	date := time.Now().AddDate(0, -6, 0)
	if start.Month() != date.Month() {
		t.Error()
	}

	if end.Month() != time.Now().Month() {
		t.Error()
	}
}

func TestPreviousSixMonths(t *testing.T) {
	start, end := previousSixMonths()

	date := time.Now().AddDate(0, -12, 0)
	if start.Month() != date.Month() {
		t.Error()
	}

	date = time.Now().AddDate(0, -6, 0)
	if end.Month() != date.Month() {
		t.Error()
	}
}

func TestThisYear(t *testing.T) {
	start, end := thisYear()
	if start.Year() != time.Now().Year() {
		t.Error()
	}

	if end.Year() != time.Now().Year() {
		t.Error()
	}
}

func TestLastYear(t *testing.T) {
	start, end := lastYear()

	date := time.Now().AddDate(-1, 0, 0)
	if start.Year() != date.Year() {
		t.Error()
	}

	if end.Year() != time.Now().Year() {
		t.Error()
	}
}

func TestAllTime(t *testing.T) {
	start, end := allTime()

	date := time.Date(epoch, 1, 1, 0, 0, 0, 0, time.Now().Location())
	if start.Year() != date.Year() {
		t.Error()
	}

	if end.Year() != time.Now().Year() {
		t.Error()
	}
}
