package main

import (
	"fmt"
	"go-chat/handlers"
	"go-chat/middleware"
	"net/http"
	"os"

	"go-chat/models"
	"go-chat/socket"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {

	// Load the .env file
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	// TODO
	// Reference the db and close connection when this function returns
	// Still not sure if this is even necessary
	db := models.GetDB()
	defer db.Close()

	// Create a new websocket hub
	wsHub := socket.NewHub()

	// Create a router
	router := mux.NewRouter()

	// CORS middleware
	router.Use(middleware.CORSHandler)
	// JWT middleware
	// router.Use(middleware.JwtAuthentication)

	// Routes
	router.HandleFunc("/user", handlers.GetUserHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/user/new", handlers.CreateUserHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", handlers.Authenticate).Methods("POST", "OPTIONS")
	// Ticketing route for ws authentication
	router.HandleFunc("/ws/auth", handlers.HandleWebSocketAuth).Methods("POST", "OPTIONS")
	// Websocket connection
	router.HandleFunc("/ws/{roomID}", wsHub.HandleWebSocketConns).Methods("GET", "POST")

	router.HandleFunc("/chat/conversations/new", handlers.CreateConversation).Methods("POST", "OPTIONS")
	router.HandleFunc("/chat/conversations", handlers.GetConversationsByUserID).Methods("GET", "OPTIONS")
	router.HandleFunc("/chat/conversations/{conversationID}", handlers.GetConversation).Methods("GET", "OPTIONS")
	router.HandleFunc("/chat/conversations/{conversationID}/messages", handlers.GetMessagesByConversationID).Methods("GET", "OPTIONS")

	// Start listening for incoming chat messages
	// go handlers.HandleWebSocketMessages()

	// Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Server
	if err := http.ListenAndServe(":"+port, router); err != nil {
		fmt.Print(err)
	}
}
