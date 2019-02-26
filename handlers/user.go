package handlers

import (
	"encoding/json"
	"fmt"
	"go-auth/models"
	u "go-auth/utils"
	"io/ioutil"
	"net/http"
)

// GetUserHandler - GET Route to for user
func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	//decode the request body into struct
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	email := user.Email

	//Create account
	resp := models.GetUserByEmail(email)
	u.Respond(w, resp)

}

// CreateUserHandler - POST route to create a user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	//decode the request body into struct
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	//Create account
	resp := user.Create()
	u.Respond(w, resp)
}

// Authenticate a user
func Authenticate(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}

	//decode the request body into struct
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)
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
