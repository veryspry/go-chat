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

// CreateConversation - POST route to create a conversation
func CreateConversation(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var users usersResp

	err := decoder.Decode(&users)

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	conv := &models.Conversation{}

	//Create new conversation
	resp := conv.Create(users.Users)

	u.Respond(w, resp)
}

// GetConversations returns all of the conversations
func GetConversations(w http.ResponseWriter, r *http.Request) {
	convs := models.GetConversations()
	u.Respond(w, convs)
}

// GetConversationsByUserID Returns all of the conversations from a single user
func GetConversationsByUserID(w http.ResponseWriter, r *http.Request) {

	idHeader := r.Header.Get("UserID")

	fmt.Println("UserID", idHeader)

	var resp map[string]interface{}

	if idHeader == "" {
		resp = u.Message(false, "Malformed userID header")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, resp)
	}

	id := u.UUIDFromString(idHeader)
	convs := models.GetConversationsByUserID(id)
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

// GetMessagesByConversationID returns all the messages from a conversation
func GetMessagesByConversationID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	convID := vars["conversationID"]

	id := u.UUIDFromString(convID)

	resp := models.GetMessagesByConversationID(id)

	u.Respond(w, resp)
}
