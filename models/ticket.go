package models

import (
	u "go-chat/utils"

	uuid "github.com/satori/go.uuid"
)

// Ticket type for db
type Ticket struct {
	BaseFields
	Token string `json:"token";sql:"-"`
}

// Create creates a new ticket
func (ticket *Ticket) Create() (map[string]interface{}, bool) {

	// Generate and set ID field using uuid v4
	id, err := uuid.NewV4()
	if err != nil {
		return u.Message(false, "Failed to create ticket, error creating ID"), false
	}
	ticket.ID = id

	db := GetDB()
	db.Create(&ticket)

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
