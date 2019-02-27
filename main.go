package main

import (
	"fmt"
	"go-auth/handlers"
	"go-auth/middleware"
	"net/http"
	"os"

	"go-auth/models"
	u "go-auth/utils"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func test(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Access-Control-Allow-Credentials", "true")
	headers.Add("Content-Type", "application/json")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")

	w.WriteHeader(http.StatusOK)

	resp := map[string]interface{}{"status": "Okay"}

	u.Respond(w, resp)
}

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
	router.Use(middleware.JwtAuthentication)

	// Routes
	router.HandleFunc("/user", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/user/new", handlers.CreateUserHandler).Methods("POST")
	router.HandleFunc("/login", handlers.Authenticate).Methods("POST", "OPTIONS")

	// Server react build
	// buildHandler := http.FileServer(http.Dir(""))
	// router.PathPrefix("/").Handler(buildHandler)

	// Serve static files
	// staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("<path to build/static>")))
	// router.PathPrefix("/static/").Handler(staticHandler)

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
