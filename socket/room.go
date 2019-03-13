package socket

import (
	"fmt"
	u "go-auth/utils"

	"github.com/satori/go.uuid"

	"github.com/gorilla/websocket"
)

// Room manages a single chat group
type Room struct {
	// id of the room which functions as a name reference
	id uuid.UUID
	// a list of clients in the room
	clients map[uuid.UUID]*Client
	// The number clients in the room
	count int
}

// Join adds a new client to a room
func (r *Room) Join(conn *websocket.Conn, userID uuid.UUID) uuid.UUID {

	r.clients[userID] = NewClient(conn)
	fmt.Printf("New client joined %s", userID)
	r.count++
	return userID
}

// Leave removes a client from a room
func (r *Room) Leave(id uuid.UUID) {
	r.count--
	delete(r.clients, id)
}

// BroadcastAll broadcasts a message to everyone in a room, including the sender
func (r *Room) BroadcastAll(senderID, roomID uuid.UUID, msg string) {
	for _, client := range r.clients {
		client.WriteMsg(senderID, roomID, msg)
	}
}

// BroadcastExc broadcasts a message to everyone, excluding the sender
func (r *Room) BroadcastExc(senderID, roomID uuid.UUID, msg string) {
	for id, client := range r.clients {
		if id != senderID {
			client.WriteMsg(senderID, roomID, msg)
		}
	}
}

// HandleMsg broadcasts a messages to a room
func (r *Room) HandleMsg(id uuid.UUID) {
	roomID := r.id
	for {
		if r.clients[id] == nil {
			break
		}
		out := <-r.clients[id].out

		if out.BroadcastAll == true {
			r.BroadcastAll(id, roomID, out.Message)
		} else {
			r.BroadcastExc(id, roomID, out.Message)
		}

	}
}

// NewRoom constructs a new room - takes in an id (from the database)
// When a new room is created, appropriate records are stored in the database and an id is passed in here
// If the room already exists, its record is retrieved and passed in here
func NewRoom(id uuid.UUID) *Room {
	r := new(Room)
	// TODO: this should be set based on a conversations id in the database
	r.id = u.NewUUID()
	r.clients = make(map[uuid.UUID]*Client)
	r.count = 0
	return r
}
