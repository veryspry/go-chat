package utils

import (
	"encoding/json"
	"net/http"
)

// Message formatter
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond - response formatter
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	// w.Header().Add("Access-Control-Allow-Origin", "*")
	// w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT")
	// w.Header().Add("Access-Control-Allow-Credentials", "true")
	// w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
