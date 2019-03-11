package socket

import (
	"log"

	"github.com/gorilla/websocket"
)

// Client is a vehicle to read and write messages from a client
type Client struct {
	// websocket connection
	conn *websocket.Conn
	// channel to pump messages through
	out chan Message
}

// NewClient constructor
func NewClient(conn *websocket.Conn) *Client {
	client := new(Client)
	client.conn = conn
	client.out = make(chan Message)
	return client
}

// WriteMsg writes a message to a client
func (c *Client) WriteMsg(msg string) {
	m := Message{}
	m.Message = msg
	err := c.conn.WriteJSON(m)
	if err != nil {
		// TODO: Update error handling to send back status, etc to the client
		log.Println("write:", err)
	}
}

// ReadLoop decodes a JSON body and pumps it out to a room channel
func (c *Client) ReadLoop() {
	defer close(c.out)

	for {
		var msg Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			// TODO: Better error handling to send back error message to the client
			log.Println("read:", err)
			break
		}
		c.out <- msg
	}
}
