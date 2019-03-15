package socket

import (
	"log"

	"go-auth/models"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// Client is a vehicle to read and write messages from a client
type Client struct {
	// websocket connection
	conn *websocket.Conn
	// channel to pump messages through
	out chan models.Message
}

// NewClient constructor
func NewClient(conn *websocket.Conn) *Client {
	client := new(Client)
	client.conn = conn
	client.out = make(chan models.Message)
	return client
}

// WriteMsg writes a message to a client
func (c *Client) WriteMsg(senderID, roomID uuid.UUID, m *models.Message) {
	// m := models.Message{}
	// m.Message = msg
	// m.UserID = senderID

	err := c.conn.WriteJSON(&m)
	if err != nil {
		// TODO: Update error handling to send back status, etc to the client
		log.Println("write:", err)
	}
}

// ReadLoop decodes a JSON body and pumps it out to a room channel
func (c *Client) ReadLoop() {
	defer close(c.out)

	for {
		var msg models.Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			// TODO: Better error handling to send back error message to the client
			log.Println("read:", err)
			break
		}
		c.out <- msg
	}
}
