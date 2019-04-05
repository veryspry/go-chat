package handlers

import (
	"encoding/json"
	"go-chat/models"
	u "go-chat/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type usersResp struct {
	UserIDs []string
}

// CreateConversation - POST route to create a conversation
// Reads the user ID from the request header and adds it to the conversation
func CreateConversation(w http.ResponseWriter, r *http.Request) {

	// Decode slice of users from request
	decoder := json.NewDecoder(r.Body)
	var users usersResp

	err := decoder.Decode(&users)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	// Add user requesting the conversation create action to the conversation
	currentUserID := r.Header.Get("userID")
	users.UserIDs = append(users.UserIDs, currentUserID)

	// Instantiate new conversation struct
	conv := &models.Conversation{}
	//Create new conversation
	resp := conv.Create(users.UserIDs)

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

	c := models.GetConversationByID(id)

	// Compose response
	resp := u.Message(false, "conversation retreived")
	// Attach conversations to the response
	resp["conversation"] = c

	u.Respond(w, resp)
}

// GetMessagesByConversationID returns all the messages from a conversation
func GetMessagesByConversationID(w http.ResponseWriter, r *http.Request) {

	// Get the conversation id from the URL path
	vars := mux.Vars(r)
	convID := vars["conversationID"]
	// Convert to type uuid.UUID
	id := u.UUIDFromString(convID)
	// Look up the conversation by ID
	conversation := models.GetConversationByID(id)

	// Attach all related messages to the conversation
	err := conversation.GetMessagesByConversationID()

	if err != nil {
		// Compose message
		resp := u.Message(true, "Error getting messages")
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, resp)
		return
	}

	// Compose response
	resp := u.Message(true, "success getting messages")
	resp["messages"] = conversation.Messages

	u.Respond(w, resp)
}
