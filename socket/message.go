package socket

// Message is a vehicle for websocket messages
type Message struct {
	// TODO add userID field
	Username string `json:"username"`
	// Message contents
	Message string `json:"message"`
	// true broadcasts to everyone, false broadcasts to all but sender
	BroadcastAll bool `json:"broadcastAll"`
}
