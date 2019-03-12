package socket

import uuid "github.com/satori/go.uuid"

// Message is a vehicle for websocket messages
type Message struct {
	UserID   uuid.UUID `json:"userID"`
	Username string    `json:"username"`
	// Message contents
	Message string `json:"message"`
	// true broadcasts to everyone, false broadcasts to all but sender
	BroadcastAll bool `json:"broadcastAll"`
}
