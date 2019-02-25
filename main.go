package main

import (
	"fmt"
	"go-auth/auth"
	"go-auth/handlers"
	"net/http"

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

	rtr := mux.NewRouter()

	rtr.HandleFunc("/user", handlers.GetUserHandler).Methods("GET")
	rtr.HandleFunc("/user", handlers.CreateUserHandler).Methods("POST")
	rtr.HandleFunc("/auth", handlers.LoginHandler).Methods("POST")
	rtr.Use(auth.JwtAuthentication)

	http.Handle("/", rtr)

	rtr.PathPrefix("../../../vs/code/ui/dist/index.html").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
