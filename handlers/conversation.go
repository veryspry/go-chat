package handlers

import (
	"encoding/json"
	"fmt"
	"go-auth/models"
	u "go-auth/utils"
	"net/http"

	"github.com/gorilla/mux"
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

// GetConversations returns all of the conversations
func GetConversations(w http.ResponseWriter, r *http.Request) {
	convs := models.GetConversations()
	u.Respond(w, convs)
}

// GetConversation returns a single conversation by ID
func GetConversation(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	convID := vars["conversationID"]

	id := u.UUIDFromString(convID)

	resp := models.GetConversationByID(id)
	u.Respond(w, resp)
}

// GetConversationMessages returns all the messages from a conversation
func GetConversationMessages(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	convID := vars["conversationID"]

	id := u.UUIDFromString(convID)

	resp := models.GetMessagesByConversationID(id)

	u.Respond(w, resp)
}
