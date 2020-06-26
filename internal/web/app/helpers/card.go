package helpers

import (
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

func PaintCard(current, previous *[]types.Summary) types.Card {
	var card types.Card
	card.Total = types.Total{Value: int64(CardGrandTotal(current))}
	card.Diff = PercentDiff(card.Total.Value, int64(CardGrandTotal(previous)))

	return card
}

func CardGrandTotal(items *[]types.Summary) float64 {
	var gt float64
	for _, item := range *items {
		gt += item.Total
	}

	return gt
}
