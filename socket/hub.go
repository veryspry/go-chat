package socket

import (
	u "go-auth/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// Hub manages all of the rooms
type Hub struct {
	// Holds a map of rooms
	hub map[uuid.UUID]*Room
	// upgrade from https connection to wss
	upgrader websocket.Upgrader
}

// HandleWebSocketConns handles websocket connections
// TODO: This needs to check for a valid "ticket" on initial upgrade
func (hub *Hub) HandleWebSocketConns(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	roomID := params["roomID"]
	roomUUID := u.UUIDFromString(roomID)

	// Upgrade initial GET request to a websocket
	ws, err := hub.upgrader.Upgrade(w, r, nil)
	if err != nil {
		msg := u.Message(false, "Websocket connection error")
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, msg)
	}

	defer ws.Close()

	room := hub.GetRoom(roomUUID)
	id := room.Join(ws)

	// Reads from the client's out bound channel and broadcasts it
	go room.HandleMsg(id)

	// Reads from client and if this loop breaks then client disconnected
	room.clients[id].ReadLoop()
	room.Leave(id)
}

// GetRoom creates a new room if it doesn't exist and returns it
func (hub *Hub) GetRoom(id uuid.UUID) *Room {

	if _, ok := hub.hub[id]; !ok {
		hub.hub[id] = NewRoom(id)
	}
	return hub.hub[id]
}

// NewHub instantiates a new instance of Hub
func NewHub() *Hub {
	hub := new(Hub)
	hub.hub = make(map[uuid.UUID]*Room)

	// Configure the upgrader
	hub.upgrader = websocket.Upgrader{
		// TODO: Update this with a better check
		// A hacky way to allow upgrade requests from any origin
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return hub
}
