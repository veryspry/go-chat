package socket

import (
	"fmt"
	"go-auth/models"

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
func (r *Room) BroadcastAll(senderID, roomID uuid.UUID, msg *models.Message) {
	for _, client := range r.clients {
		client.WriteMsg(senderID, roomID, msg)
	}
}

// BroadcastExc broadcasts a message to everyone, excluding the sender
func (r *Room) BroadcastExc(senderID, roomID uuid.UUID, msg *models.Message) {
	for id, client := range r.clients {
		if id != senderID {
			client.WriteMsg(senderID, roomID, msg)
		}
	}
}

// HandleMsg broadcasts a messages to a room and saves it to the db
func (r *Room) HandleMsg(id uuid.UUID) {

	for {
		if r.clients[id] == nil {
			break
		}
		out := <-r.clients[id].out

		// Save the message to the db
		m := models.Message{}
		m.Message = out.Message
		m.UserID = id
		roomID := r.id

		// Ignore any empty message bodys
		if m.Message != "" {
			_ = m.Create(id, roomID)
		}

		// if resp["message"] != "success" {
		// 	// TODO: Update error handling to send back status, etc to the client
		// 	log.Println("save:", resp)
		// }

		if out.BroadcastAll == true {
			r.BroadcastAll(id, roomID, &m)
		} else {
			r.BroadcastExc(id, roomID, &m)
		}

	}
}

// NewRoom constructs a new room
// When a new room is created, appropriate records are stored in the database and an id is passed in here
// If the room already exists, its record is retrieved and passed in here
func NewRoom(id uuid.UUID) *Room {
	r := new(Room)
	// TODO: this should be set based on a conversations id in the database
	// r.id = u.NewUUID()
	r.id = id
	r.clients = make(map[uuid.UUID]*Client)
	r.count = 0
	return r
}
