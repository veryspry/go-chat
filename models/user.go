package models

import (
	"net/http"
	"os"
	"strings"

	u "go-auth/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User type for db
type User struct {
	BaseFields
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	Token        string `json:"token" sql:"-"`
	Conversation []*Conversation
}

//Validate incoming user details
func (user *User) Validate() (map[string]interface{}, bool) {

	// TODO - BETTER EMAIL VALIDATION
	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	// TODO - PASSWORD VALIDATION

	//Email must be unique
	temp := &User{}

	//check for errors and duplicate emails
	err := GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

// Create a new user
func (user *User) Create() map[string]interface{} {

	// Validate the request
	if resp, ok := user.Validate(); !ok {
		return resp
	}

	// Hash and set password on User
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	// Generate and set ID field using uuid v4
	id, err := uuid.NewV4()
	if err != nil {
		return u.Message(false, "Failed to create account, error creating ID")
	}
	user.ID = id

	// Create new JWT token for the newly registered user
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	db := GetDB()
	db.Create(&user)

	// Delete password
	user.Password = ""

	// Compose a response
	response := u.Message(true, "user has been created")
	// Attach the user to the response
	response["user"] = user

	// TODO: this should be read from the user, but can't figure out how to do that yet
	// Attach the token
	response["token"] = tokenString
	return response
}

// Login a user
func Login(email, password string, w http.ResponseWriter) map[string]interface{} {

	user := &User{}

	// Look up user record
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	// Check if passwords match
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Password invalid")
	}

	// Delete the password
	user.Password = ""

	// Create JWT token and store it in response
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	// Compose response message and attach user to the response message
	resp := u.Message(true, "Logged in")
	resp["user"] = user
	resp["token"] = tokenString

	return resp
}

// GetUserByEmail - return a single user by email address
func GetUserByEmail(email string) map[string]interface{} {

	// Get the db connection
	db := GetDB()
	user := &User{}

	// Lookup the user
	db.Table("users").Where("email = ?", email).First(user)
	// If user doesn't exist
	if user.Email == "" {
		return nil
	}

	// Compose response message
	resp := u.Message(true, "User found")
	// Add user to the response
	resp["user"] = user

	return resp
}

// GetUserByID - return a single user by email address
func GetUserByID(id uuid.UUID) *User {

	// Get the db connection
	db := GetDB()
	user := &User{}

	// Lookup the user
	db.Table("users").Where("ID = ?", id).First(user)
	// If user doesn't exist
	if user.Email == "" {
		return nil
	}

	return user
}

// GetUserByToken returns a user by their token
func GetUserByToken(token string) map[string]interface{} {

	// Get the db connection
	db := GetDB()
	user := &User{}

	// Lookup the ticket
	db.Table("users").Where("token = ?", token).First(user)
	// If ticket doesn't exist
	if user.Token == "" {
		return nil
	}

	// Compose response message
	resp := u.Message(true, "Ticket found")
	// Add ticket to the response
	resp["user"] = user
	return resp
}
