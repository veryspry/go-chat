package models

import (
	"fmt"
	u "go-auth/utils"

	uuid "github.com/satori/go.uuid"
)

// Conversation is what groups together users into a single conversation thread
type Conversation struct {
	BaseFields
	Messages []*Message
	Users    []*User `gorm:"many2many:user_conversations"`
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
		// TODO: Will this nested _ cause issue?
		users = append(users, user)
	}

	c.Users = users

	fmt.Println("USERS: ", users)
	fmt.Println("CONVERSATION: ", &c.Users)

	db := GetDB()
	db.Create(&c)

	// Compose a response
	response := u.Message(false, "conversation has been created")
	// Attach the user to the response
	response["conversation"] = c

	return response
}
