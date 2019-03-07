package main

import (
	"fmt"
	"go-auth/handlers"
	"go-auth/middleware"
	"net/http"
	"os"

	"go-auth/models"

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

	// Create a router
	router := mux.NewRouter()

	// CORS middleware
	router.Use(middleware.CORSHandler)
	// JWT middleware
	// router.Use(middleware.JwtAuthentication)

	// Routes
	router.HandleFunc("/user", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/user/new", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", handlers.Authenticate).Methods("POST", "OPTIONS")
	router.HandleFunc("/ws", handlers.HandleWebSocketConns)

	// Start listening for incoming chat messages
	go handlers.HandleWebSocketMessages()

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
