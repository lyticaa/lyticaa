package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type CustomTransaction struct {
	gorm.Model
	Id                     int64
	DateTime               time.Time
	SettlementId           int64
	OrderType              OrderType `gorm:"foreignkey:Id"`
	OrderId                string
	Sku                    string
	Quantity               int64
	Marketplace            Marketplace        `gorm:"foreignkey:Id"`
	Fulfillment            Fulfillment        `gorm:"foreignkey:Id"`
	TaxCollectionModel     TaxCollectionModel `gorm:"foreignkey:Id"`
	ProductSales           float64
	ProductSalesTax        float64
	ShippingCredits        float64
	ShippingCreditsTax     float64
	GiftWrapCredits        float64
	GiftWrapCreditsTax     float64
	PromotionalRebates     float64
	PromotionalRebatesTax  float64
	MarketplaceWithheldTax float64
	SellingFees            float64
	FBAFees                float64
	OtherTransactionFees   float64
	Order                  float64
	Total                  float64
}

// TODO : Support for ONCONFLICT/UPSERT
func SaveCustomTransaction(txn CustomTransaction, db *gorm.DB) {
	//if txn.Id > 0 {
	//	UpdateCustomTransaction(txn, db)
	//} else {
	//	db.Create(txn)
	//}
}

func UpdateCustomTransaction(txn CustomTransaction, db *gorm.DB) {
	db.Model(&txn).Update("updated_at", time.Now())
}
