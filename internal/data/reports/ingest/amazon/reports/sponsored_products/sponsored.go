package sponsored_products

import (
	"github.com/jmoiron/sqlx"
)

type SponsoredProducts struct {
	db *sqlx.DB
}

func NewSponsoredProducts(db *sqlx.DB) *SponsoredProducts {
	return &SponsoredProducts{
		db: db,
	}
}
