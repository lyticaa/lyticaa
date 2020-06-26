package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Marketplace struct {
	Id        int64
	Name      string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (m *Marketplace) Load(db *sqlx.DB) *[]Marketplace {
	var marketplaces []Marketplace

	query := `SELECT id,name,created_at,updated_at FROM marketplaces ORDER BY id DESC`
	_ = db.Select(&marketplaces, query)

	return &marketplaces
}
