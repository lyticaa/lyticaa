package custom_transactions

import (
	"strconv"
	"strings"
	"time"
)

func (c *CustomTransactions) stripComma(data string) string {
	return strings.Replace(data, ",", "", -1)
}

func (c *CustomTransactions) amazonTransactionType(row map[string]string) int64 {
	transactionType, _ := c.amazonTransactionTypeIDByName(row["type"])
	return transactionType
}

func (c *CustomTransactions) orderID(row map[string]string) string {
	return row["order id"]
}

func (c *CustomTransactions) sku(row map[string]string) string {
	return row["sku"]
}

func (c *CustomTransactions) amazonMarketplace(row map[string]string) int64 {
	marketplace, _ := c.amazonMarketplaceIDByName(strings.ToLower(row["marketplace"]))
	return marketplace
}

func (c *CustomTransactions) amazonFulfillment(row map[string]string) int64 {
	fulfillment, _ := c.amazonFulfillmentIDByName(row["fulfillment"])
	return fulfillment
}

func (c *CustomTransactions) amazonTaxCollectionModel(row map[string]string) int64 {
	taxCollectionModel, _ := c.amazonTaxCollectionModelIDByName(row["tax collection model"])
	return taxCollectionModel
}

func (c *CustomTransactions) dateTime(row map[string]string) time.Time {
	dt, _ := time.Parse("Jan 2, 2006 3:04:05 PM MST", row["date/time"])
	return dt
}

func (c *CustomTransactions) settlementID(row map[string]string) int64 {
	settlementID, _ := strconv.ParseInt(row["settlement id"], 10, 64)
	return settlementID
}

func (c *CustomTransactions) quantity(row map[string]string) int64 {
	quantity, _ := strconv.ParseInt(row["quantity"], 10, 64)
	return quantity
}

func (c *CustomTransactions) productSales(row map[string]string) float64 {
	productSales, _ := strconv.ParseFloat(c.stripComma(row["product sales"]), 64)
	return productSales
}

func (c *CustomTransactions) productSalesTax(row map[string]string) float64 {
	productSalesTax, _ := strconv.ParseFloat(c.stripComma(row["product sales tax"]), 64)
	return productSalesTax
}

func (c *CustomTransactions) shippingCredits(row map[string]string) float64 {
	shippingCredits, _ := strconv.ParseFloat(c.stripComma(row["shipping credits"]), 64)
	return shippingCredits
}

func (c *CustomTransactions) shippingCreditsTax(row map[string]string) float64 {
	shippingCreditsTax, _ := strconv.ParseFloat(c.stripComma(row["shipping credits tax"]), 64)
	return shippingCreditsTax
}

func (c *CustomTransactions) giftwrapCredits(row map[string]string) float64 {
	giftwrapCredits, _ := strconv.ParseFloat(c.stripComma(row["gift wrap credits"]), 64)
	return giftwrapCredits
}

func (c *CustomTransactions) giftwrapCreditsTax(row map[string]string) float64 {
	giftwrapCreditsTax, _ := strconv.ParseFloat(c.stripComma(row["giftwrap credits tax"]), 64)
	return giftwrapCreditsTax
}

func (c *CustomTransactions) promotionalRebates(row map[string]string) float64 {
	promotionalRebates, _ := strconv.ParseFloat(c.stripComma(row["promotional rebates"]), 64)
	return promotionalRebates
}

func (c *CustomTransactions) promotionalRebatesTax(row map[string]string) float64 {
	promotionalRebatesTax, _ := strconv.ParseFloat(c.stripComma(row["promotional rebates tax"]), 64)
	return promotionalRebatesTax
}

func (c *CustomTransactions) marketplaceWithheldTax(row map[string]string) float64 {
	marketplaceWithheldTax, _ := strconv.ParseFloat(c.stripComma(row["marketplace withheld tax"]), 64)
	return marketplaceWithheldTax
}

func (c *CustomTransactions) sellingFees(row map[string]string) float64 {
	sellingFees, _ := strconv.ParseFloat(c.stripComma(row["selling fees"]), 64)
	return sellingFees
}

func (c *CustomTransactions) fbaFees(row map[string]string) float64 {
	fbaFees, _ := strconv.ParseFloat(c.stripComma(row["fba fees"]), 64)
	return fbaFees
}

func (c *CustomTransactions) otherTransactionFees(row map[string]string) float64 {
	otherTransactionFees, _ := strconv.ParseFloat(c.stripComma(row["other transaction fees"]), 64)
	return otherTransactionFees
}

func (c *CustomTransactions) other(row map[string]string) float64 {
	other, _ := strconv.ParseFloat(c.stripComma(row["other"]), 64)
	return other
}

func (c *CustomTransactions) total(row map[string]string) float64 {
	total, _ := strconv.ParseFloat(c.stripComma(row["total"]), 64)
	return total
}
