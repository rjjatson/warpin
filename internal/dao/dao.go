package dao

// NotifMessage DAO of notif message
type NotifMessage struct {
	storage []string
}

// New creates new notif message
func New() *NotifMessage {
	return &NotifMessage{
		storage: make([]string, 0),
	}
}

// GetAll gets all stored message
func (n *NotifMessage) GetAll() ([]string, error) {
	return n.storage, nil
}

// Store stores a notif message
func (n *NotifMessage) Store(msg string) error {
	n.storage = append(n.storage, msg)
	return nil
}
