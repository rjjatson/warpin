package model

import "time"

// SendNotifRequest is the request model for send notif API
type SendNotifRequest struct {
	Message string `json:"message"`
}

// GetAllNotifResponse is the request model for send notif API
type GetAllNotifResponse struct {
	Notifications []Notification `json:"notifications"`
}

// Notification is the object model for API response
type Notification struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

// NotificationStore is the object model for stored notification
type NotificationStore struct {
	Message string
	Time    time.Time
}
