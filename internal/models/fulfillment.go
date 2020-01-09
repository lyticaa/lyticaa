package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Fulfillment struct {
	Id        int64
	Name      string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func GetFulfillments(db *sqlx.DB) []Fulfillment {
	var fulfillments []Fulfillment
	_ = db.Select(&fulfillments, "SELECT id,name,created_at,updated_at FROM fulfillments ORDER BY id DESC")

	return fulfillments
}
