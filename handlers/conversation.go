package handlers

import (
	"encoding/json"
	"fmt"
	"go-auth/models"
	u "go-auth/utils"
	"net/http"
)

type usersResp struct {
	Users []string
}

// CreateConversationHandler - POST route to create a user
func CreateConversationHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var users usersResp

	err := decoder.Decode(&users)

	fmt.Println("\n Response: \n", users.Users)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	// fmt.Println("\n Conv: \n", conv)

	conv := new(models.Conversation)

	//Create new conversation
	resp := conv.Create(users.Users)

	u.Respond(w, resp)
}
