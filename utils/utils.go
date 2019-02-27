package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Message formatter
func Message(isAuthed bool, message string) map[string]interface{} {
	return map[string]interface{}{"isAuthenticated": isAuthed, "message": message}
}

// Respond - response formatter
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	json.NewEncoder(w).Encode(data)
}

// CreateSession - create a session in th Redis store and give the user a session cookie
func CreateSession(w http.ResponseWriter, token, email string) error {

	// Get the redis cache
	// cache := redis.GetRedisCache()
	// Set the token in the cache, along with the user whom it represents
	// The token has an expiry time of 3600 seconds (one hour)
	_, err := cache.Do("SETEX", token, "3600", email)
	if err != nil {
		return err
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds, the same as the cache
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   token,
		Expires: time.Now().Add(3600 * time.Second),
	})

	return nil

}

var cache redis.Conn

func initCache() {
	// Initialize the redis connection to a redis instance running on your local machine
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	// Assign the connection to the package level `cache` variable
	cache = conn
}

func main() {
	initCache()
}
