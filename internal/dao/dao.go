package dao

import (
	"time"
)

// Notification is the object model for stored notification
type Notification struct {
	Message string
	Time    time.Time
}

// NotifMessage DAO of notif message
type NotifMessage struct {
	storage []Notification
}

// New creates new notif message
func New() *NotifMessage {
	return &NotifMessage{
		storage: make([]Notification, 0),
	}
}

// GetAll gets all stored message
func (n *NotifMessage) GetAll() ([]Notification, error) {
	return n.storage, nil
}

// Store stores a notif message
func (n *NotifMessage) Store(msg string, ts time.Time) error {
	n.storage = append(n.storage, Notification{
		Message: msg,
		Time:    ts,
	})
	return nil
}
