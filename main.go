package main

import (
	"go-auth/handlers"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// http://127.0.0.1:53088/?key=5dd40369-e614-433c-8aad-c516d1933d50

func main() {

	rtr := mux.NewRouter()
	// rtr.HandleFunc("/user", queries.CreateUser).Methods("POST")

	rtr.HandleFunc("/user", handlers.GetUserHandler).Methods("GET")
	rtr.HandleFunc("/user", handlers.CreateUserHandler).Methods("POST")

	http.Handle("/", rtr)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
