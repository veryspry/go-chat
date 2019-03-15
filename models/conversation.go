package models

import (
	"fmt"
	u "go-auth/utils"

	uuid "github.com/satori/go.uuid"
)

// Conversation is what groups together users into a single conversation thread
type Conversation struct {
	BaseFields
	// gorm.Model
	Messages []*Message
	Users    []*User `gorm:"many2many:user_conversation_join;"`
	// association_foreignkey:userId;foreignkey:conversationId
}

// Create creates a new conversation that has all of the passed in users
func (c *Conversation) Create(usrIds []string) map[string]interface{} {

	// Generate and set ID field using uuid v4
	id, err := uuid.NewV4()
	if err != nil {
		return u.Message(false, "Failed to create account, error creating ID")
	}
	c.ID = id

	var users []*User
	// var id uuid.UUID

	for _, usrID := range usrIds {
		id := u.UUIDFromString(usrID)
		user := GetUserByID(id)
		users = append(users, user)
	}

	c.Users = users

	fmt.Println("USERS: ", &users)
	fmt.Println("CONVERSATION: ", c.Users)

	db := GetDB()
	db.Debug().Create(&c)

	// Compose a response
	response := u.Message(false, "conversation has been created")
	// Attach the user to the response
	response["conversation"] = c

	return response
}

// GetConversationMessages returns the messages from a conversation
func (c *Conversation) GetConversationMessages() map[string]interface{} {

	message := new(Message)

	db := GetDB()
	m := db.Debug().Model(&c).Related(&message)

	// Compose response
	response := u.Message(false, "messages retreived")
	// Attach messages to the response
	response["messages"] = m

	return response

}

// GetConversations returns all conversations from the db
func GetConversations() map[string]interface{} {

	var conv []*Conversation

	db := GetDB()
	convs := db.Find(&conv)

	// Compose response
	resp := u.Message(false, "conversations retreived")
	// Attatch conversations to the response
	resp["conversations"] = convs

	return resp
}

// GetConversationByID gets a single conversation by id
func GetConversationByID(id uuid.UUID) map[string]interface{} {

	c := Conversation{}
	db := GetDB()
	conv := db.Where("id = ?", id).Find(&c)

	// Compose response
	resp := u.Message(false, "conversation retreived")
	// Attach conversations to the response
	resp["conversation"] = conv

	return resp
}
