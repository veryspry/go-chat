package models

import (
	u "go-auth/utils"

	uuid "github.com/satori/go.uuid"
)

// Message is a single message that belongs to a user and a conversation
type Message struct {
	BaseFields
	UserID         uuid.UUID `json:"userID"`
	User           *User
	ConversationID uuid.UUID `json:"roomID"`
	Conversation   *Conversation
	Message        string `json:"message"`
	// true broadcasts to everyone, false broadcasts to all but sender
	BroadcastAll bool `json:"broadcastAll"`
}

// Create saves a new Message to the db
func (m *Message) Create(senderID, roomID uuid.UUID) map[string]interface{} {

	// Generate and set ID field using uuid v4
	id, err := uuid.NewV4()
	if err != nil {
		return u.Message(false, "Failed to create account, error creating ID")
	}
	m.ID = id
	m.ConversationID = roomID
	m.UserID = senderID

	db := GetDB()
	db.Create(&m)

	// Compose a response
	response := u.Message(false, "success")
	// Attach the user to the response
	response["message"] = m

	return response
}
