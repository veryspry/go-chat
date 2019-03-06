package handlers

import (
	"encoding/json"
	"fmt"
	"go-auth/models"
	u "go-auth/utils"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/securecookie"
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

	resp := models.Login(user.Email, user.Password, w)

	token := resp["token"].(string)

	// Create a session and session cookie

	// Get token_secret
	tokenSecret := os.Getenv("token_secret")
	// Get db uri
	dbURI := models.GetDBURI()

	// Fetch new store.
	store, err := pgstore.NewPGStore(dbURI, []byte(tokenSecret))
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer store.Close()

	// Run a background goroutine to clean up expired sessions from the database
	defer store.StopCleanup(store.Cleanup(time.Minute * 5))

	// Create a session.
	session, err := store.New(r, token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, "Internal server error"))
	}

	// Add a value
	session.Values["userEmail"] = user.Email
	fmt.Print(session.Options.Path)

	// Save
	if err = store.Save(r, w, session); err != nil {
		log.Fatalf("Error saving session: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		u.Respond(w, u.Message(false, "Internal server error"))
	}

	// if localhost, omit options.domain
	// Essentially, we need to do this because most browsers won't allow you to set a cookie if the domain field on cookie is present and you are on requeting from localhost
	referer := r.Referer()

	if strings.Contains(referer, "dev") {
		fmt.Print("\n", "REFERER: ", r.Referer())
		// Keep the session ID key in a cookie so it can be looked up in DB later.
		encoded, err := securecookie.EncodeMulti(session.Name(), session.ID, store.Codecs...)
		if err != nil {
			msg := u.Message(false, "Internal server error")
			w.WriteHeader(http.StatusInternalServerError)
			u.Respond(w, msg)
		}

		cookie := session.Name() + "=" + encoded

		w.Header().Set("Set-Cookie", cookie)
	}

	fmt.Print("\n", "Headers: ", w.Header())

	u.Respond(w, resp)
}
