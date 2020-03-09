package dao

import (
	"time"
	"warpin/internal/model"
)

// NotificationDAO interface for notification store
type NotificationDAO interface {
	GetAll() ([]model.NotificationStore, error)
	Store(msg string, ts time.Time) error
}
