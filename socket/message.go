package socket

import uuid "github.com/satori/go.uuid"

// MessageTest is a vehicle for websocket messages
type MessageTest struct {
	UserID   uuid.UUID `json:"userID"`
	Username string    `json:"username"`
	// Message contents
	Message string `json:"message"`
	// true broadcasts to everyone, false broadcasts to all but sender
	BroadcastAll bool `json:"broadcastAll"`
}
