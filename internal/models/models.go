package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Filter struct {
	DateRange string
	Start     int64
	Length    int64
	Sort      int64
	Dir       string
	StartDate time.Time
	EndDate   time.Time
}

const (
	TotalSales         = "total_sales"
	UnitsSold          = "units_sold"
	AmazonCosts        = "amazon_costs"
	ProductCosts       = "product_costs"
	AdvertisingSpend   = "advertising_spend"
	Refunds            = "refunds"
	ShippingCredits    = "shipping_credits"
	PromotionalRebates = "promotional_rebates"
	TotalCosts         = "total_costs"
	GrossMargin        = "gross_margin"
	NetMargin          = "net_margin"
)

type Model interface {
	FetchOne(context.Context, *sqlx.DB) interface{}
	FetchBy(context.Context, *sqlx.DB) interface{}
	FetchAll(context.Context, map[string]interface{}, *Filter, *sqlx.DB) interface{}
	Count(context.Context, map[string]interface{}, *sqlx.DB) int64
	Create(context.Context, *sqlx.DB) error
	Update(context.Context, *sqlx.DB) error
	Delete(context.Context, *sqlx.DB) error
}

func NewFilter() *Filter {
	return &Filter{}
}

func OrderBy(mapping map[int64]string, filter *Filter) string {
	return fmt.Sprintf("%v %v", sortColumn(mapping, filter.Sort), filter.Dir)
}

func sortColumn(columnMap map[int64]string, columnIDx int64) string {
	if columnIDx > int64(len(columnMap)) {
		return columnMap[0]
	}

	return columnMap[columnIDx]
}
