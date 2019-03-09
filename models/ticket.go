package models

import (
	u "go-auth/utils"

	"github.com/jinzhu/gorm"
)

// Ticket type for db
type Ticket struct {
	gorm.Model
	Token string `json:"token";sql:"-"`
}

// Create creates a new ticket
func (ticket *Ticket) Create() (map[string]interface{}, bool) {

	db := GetDB()
	db.Create(&ticket)

	if ticket.ID <= 0 {
		return u.Message(false, "Failed to authenticate websocket connection"), false
	}

	// Compose a response
	response := u.Message(true, "user has been created")
	// Attach the user to the response
	response["ticket"] = ticket

	return response, true
}

// GetTicketByToken looks up a ticket by token
func GetTicketByToken(token string) map[string]interface{} {
	// Get the db connection
	db := GetDB()
	ticket := &Ticket{}

	// Lookup the ticket
	db.Table("tickets").Where("token = ?", token).First(ticket)
	// If ticket doesn't exist
	if ticket.Token == "" {
		return nil
	}

	// Compose response message
	resp := u.Message(true, "Ticket found")
	// Add ticket to the response
	resp["ticket"] = ticket
	return resp
}
