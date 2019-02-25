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

	user, err := models.GetUserByEmail(email)
	if err != nil {
		panic("Error getting user")
	}

	fmt.Println(user)

	w.Write([]byte("User found"))

}

// CreateUserHandler - POST route to create a user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	message := "Success"

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		message = "error reading request body"
		w.Write([]byte(message))
	}
	// Unmarshal the json
	byt := []byte(body)
	var dat map[string]string
	if err := json.Unmarshal(byt, &dat); err != nil {
		message = "error reading request body"
		w.Write([]byte(message))
		panic(err)
	}

	email := dat["email"]
	password := dat["password"]

	usr := models.User{Email: email, Password: password}

	_, err = usr.Create()

	if err != nil {
		message = "error creating user"
		w.Write([]byte(message))
		return
	}

	w.Write([]byte(message))
}

// LoginHandler login handler
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	errMsg := "error reading request body"

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(errMsg))
	}

	// Unmarshal the json
	byt := []byte(body)
	var dat map[string]string
	if err := json.Unmarshal(byt, &dat); err != nil {
		w.Write([]byte(errMsg))
		panic(err)
	}

	email := dat["email"]
	password := dat["password"]

	resp := models.Login(email, password)
	fmt.Println(resp)

	w.Write([]byte("success"))
}
