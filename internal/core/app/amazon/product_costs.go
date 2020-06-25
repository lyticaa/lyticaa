package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

// ((cost of product) * quantity sold) + ((cost of product) * quantity sold as coupons) + SKU ads + refunds

func (a *Amazon) ProductCosts(txns *[]models.Transaction) []types.Summary {
	var productCosts []types.Summary
	return productCosts
}
