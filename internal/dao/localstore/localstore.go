package localstore

import (
	"time"
	"warpin/internal/model"
)

// NotifMessage DAO of notif message
type NotifMessage struct {
	storage []model.NotificationStore
}

// New creates new notif message
func New() *NotifMessage {
	return &NotifMessage{
		storage: make([]model.NotificationStore, 0),
	}
}

// GetAll gets all stored message
func (n *NotifMessage) GetAll() ([]model.NotificationStore, error) {
	return n.storage, nil
}

// Store stores a notif message
func (n *NotifMessage) Store(msg string, ts time.Time) error {
	n.storage = append(n.storage, model.NotificationStore{
		Message: msg,
		Time:    ts,
	})
	return nil
}
