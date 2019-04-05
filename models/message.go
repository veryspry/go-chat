package models

import (
	"fmt"
	u "go-chat/utils"

	uuid "github.com/satori/go.uuid"
)

// Message is a single message that belongs to a user and a conversation
type Message struct {
	BaseFields
	// Message belongs one User
	UserID uuid.UUID `json:"userID"`
	User   User      `json:"user"`
	// Message belongs one Conversation
	ConversationID uuid.UUID    `json:"roomID"`
	Conversation   Conversation `json:"conversation"`
	Message        string       `json:"message"`
	// true broadcasts to everyone, false broadcasts to all but sender
	BroadcastAll bool `json:"broadcastAll"`
}

// Create saves a new Message to the db
func (m *Message) Create(senderID, roomID uuid.UUID) map[string]interface{} {

	fmt.Println("CREATING MESSAGE: ", m.Message)

	// Generate and set ID field using uuid v4
	id, err := uuid.NewV4()
	if err != nil {
		return u.Message(false, "Failed to save message, error creating ID")
	}
	m.ID = id

	m.ConversationID = roomID
	m.UserID = senderID

	db := GetDB()
	db.Create(&m)

	// Compose a response
	response := u.Message(false, "message has been created")
	// Attach the user to the response
	response["message"] = m
	return response
}

// GetMessagesByConversationID returns all messages from a conversation when given a conversation ID
func GetMessagesByConversationID(id uuid.UUID) map[string]interface{} {

	var m []*Message

	db := GetDB()
	messages := db.Order("created_at asc").Where("conversation_id = ?", id).Find(&m)

	resp := u.Message(false, "messages retrieved")
	resp["messages"] = messages
	return resp
}
