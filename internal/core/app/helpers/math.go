package helpers

import (
	"math"
)

func PercentDiff(current, previous int64) float64 {
	if current == 0 && previous == 0 {
		return 0.0
	}

	return math.Round(float64(current-previous) / (float64(current+previous) / 2) * 100)
}
