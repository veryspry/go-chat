package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Message formatter
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond - response formatter
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	fmt.Print("data", data)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
