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

func LoadNotificationsByUser(userId int64, filter *Filter, db *sqlx.DB) *[]Notification {
	var notifications []Notification

	query := `SELECT notification, created_at FROM notifications WHERE user_id = $1 AND created_at BETWEEN $2 AND $3 ORDER BY $4 LIMIT $5 OFFSET $6`
	err := db.Select(
		&notifications,
		query,
		userId,
		filter.StartDate,
		filter.EndDate,
		fmt.Sprintf("%v %v", sortColumn(notificationSortMap, filter.Sort), filter.Dir),
		filter.Length,
		filter.Start,
	)

	if err != nil {
		return &[]Notification{}
	}

	if len(notifications) == 0 {
		return &[]Notification{}
	}

	return &notifications
}

func TotalNotificationsByUser(userId int64, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM notifications WHERE user_id = $1`
	_ = db.QueryRow(query, userId).Scan(&count)

	return count
}

func (n *Notification) Save(db *sqlx.DB) error {
	query := `UPDATE notifications SET notification = :notification, updated_at = :updated_at WHERE id = :id`
	_, err := db.NamedExec(query,
		map[string]interface{}{
			"user_id":      n.UserId,
			"notification": n.Notification,
			"created_at":   time.Now(),
			"updated_at":   time.Now(),
			"id":           n.Id,
		})

	if err != nil {
		return err
	}

	return nil
}
