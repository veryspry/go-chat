package handlers

import (
	u "go-auth/utils"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	// TODO: Update this with a better check
	// A hacky way to allow upgrade requests from any origin
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Message is a structure for a websocket message
type Message struct {
	// TODO add userID field
	Username string `json:"username"`
	Message  string `json:"message"`
}

// HandleWebSocketConns handles websocket connections
func HandleWebSocketConns(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		msg := u.Message(false, "Websocket connection error")
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, msg)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()
	// Register our new client
	clients[ws] = true
	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			msg := u.Message(false, "Websocket connection error")
			w.WriteHeader(http.StatusInternalServerError)
			u.Respond(w, msg)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

// HandleWebSocketMessages pushes out messsages and removes clients if they have closed a connection
func HandleWebSocketMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
