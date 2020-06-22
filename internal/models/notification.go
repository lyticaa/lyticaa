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

func CreateNotification(userId int64, notification string, db *sqlx.DB) error {
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
			"user_id":      userId,
			"notification": notification,
			"created_at":   time.Now(),
			"updated_at":   time.Now(),
		})

	if err != nil {
		return err
	}

	return nil
}

func FindNotificationsByUser(userId int64, filter *Filter, db *sqlx.DB) *[]Notification {
	var notifications []Notification

	err := db.Select(
		&notifications,
		`SELECT notification, created_at FROM notifications WHERE user_id = $1 AND created_at BETWEEN $2 AND $3 ORDER BY $4 LIMIT $5 OFFSET $6`,
		userId,
		filter.StartDate,
		filter.EndDate,
		fmt.Sprintf("%v %v", sortColumn(notificationSortMap, filter.Sort), filter.Dir),
		filter.Length,
		filter.Start,
	)

	if err != nil {
		logger().Error().Err(err).Msgf("unable to load the notifications for the user %v", userId)
		return &[]Notification{}
	}

	if len(notifications) == 0 {
		return &[]Notification{}
	}

	return &notifications
}

func TotalNotificationsByUser(userId int64, db *sqlx.DB) int64 {
	var count int64

	err := db.QueryRow(`SELECT COUNT(id) FROM notifications WHERE user_id = $1`, userId).Scan(&count)
	if err != nil {
		logger().Error().Err(err).Msgf("unable to count the notifications for the user %v", userId)
	}

	return count
}
