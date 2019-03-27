package utils

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/satori/go.uuid"
)

// Message formatter
func Message(isAuthed bool, message string) map[string]interface{} {
	return map[string]interface{}{"isAuthenticated": isAuthed, "message": message}
}

// Respond - response formatter
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	json.NewEncoder(w).Encode(data)
}

// NewUUID returns a new uuid version 4
func NewUUID() uuid.UUID {
	id, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error creating UUID")
	}
	return id
}

// UUIDFromString parses a uuid from a string
func UUIDFromString(id string) uuid.UUID {
	uuid, err := uuid.FromString(id)
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	return uuid
}
