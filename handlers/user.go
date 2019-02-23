package handlers

import (
	"encoding/json"
	"fmt"
	"go-auth/models"
	"io/ioutil"
	"net/http"
)

// GetUserHandler - GET Route to for user
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	// Unmarshal the json
	byt := []byte(body)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	email := dat["email"].(string)

	fmt.Println(email)

	user, err := models.User.Create()
	if err != nil {
		panic("Error getting user")
	}

	fmt.Println(user)

	w.Write([]byte("User found"))

}

// CreateUserHandler - POST route to create a user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	// Unmarshal the json
	byt := []byte(body)
	var dat map[string]string
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}

	email := dat["email"]
	password := dat["password"]

	usr := map[string]interface{}{"email": email, "password": password}

	err = queries.CreateUser(*usr)

	message := "Success"

	if err != nil {
		message = "Error creating account"
	}

	w.Write([]byte(message))
}
