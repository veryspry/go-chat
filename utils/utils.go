package utils

import (
	"encoding/json"
	"net/http"
)

// Message formatter
func Message(isAuthed bool, message string) map[string]interface{} {
	return map[string]interface{}{"isAuthenticated": isAuthed, "message": message}
}

// Respond - response formatter
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	json.NewEncoder(w).Encode(data)
}
