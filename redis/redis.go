package redis

import (
	"github.com/gomodule/redigo/redis"
)

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

// GetRedisCache Return ref to redis store
func GetRedisCache() redis.Conn {
	return cache
}

func main() {
	initCache()
}
