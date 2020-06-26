package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Notification struct {
	Id           int64     `db:"id"`
	UserId       int64     `db:"user_id"`
	Notification string    `db:"notification"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

var (
	notificationSortMap = map[int64]string{
		0: "notification",
		1: "created_at",
	}
)

func (n *Notification) Load(filter *Filter, db *sqlx.DB) *[]Notification {
	var notifications []Notification

	query := `SELECT notification, created_at FROM notifications WHERE user_id = $1 AND created_at BETWEEN $2 AND $3 ORDER BY $4 LIMIT $5 OFFSET $6`
	_ = db.Select(
		&notifications,
		query,
		n.UserId,
		filter.StartDate,
		filter.EndDate,
		fmt.Sprintf("%v %v", sortColumn(notificationSortMap, filter.Sort), filter.Dir),
		filter.Length,
		filter.Start,
	)

	return &notifications
}

func (n *Notification) Total(db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM notifications WHERE user_id = $1`
	_ = db.QueryRow(query, n.UserId).Scan(&count)

	return count
}

func (n *Notification) Save(db *sqlx.DB) error {
	query := `INSERT INTO notifications (
					user_id,
					notification,
					created_at,
					updated_at)
				VALUES (
					:user_id,
					:notification,
					:created_at,
					:updated_at)`

	_, err := db.NamedExec(query,
		map[string]interface{}{
			"user_id":      n.UserId,
			"notification": n.Notification,
			"created_at":   time.Now(),
			"updated_at":   time.Now(),
		})

	if err != nil {
		return err
	}

	return nil
}
