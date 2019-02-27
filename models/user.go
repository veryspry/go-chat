package models

import (
	"net/http"
	"os"
	"strings"

	u "go-auth/utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User type for db
type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Token    string `json:"token";sql:"-"`
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

	// TODO debug why a uuid won't save to the db
	// Generate and set ID field using uuid v4
	// id, err := uuid.NewV4()
	// if err != nil {
	// 	u.Message(false, "Failed to create account, error creating ID")
	// }
	// user.ID = id

	// Create new JWT token for the newly registered user
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	db := GetDB()
	db.Create(&user)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

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
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// Passwords don't match
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Password invalid")
	}

	// Delete the password
	user.Password = ""

	// Create JWT token
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	// Store the token in the response
	user.Token = tokenString

	// Create a session"
	err = u.CreateSession(w, tokenString, user.Email)
	if err != nil {
		// If there is an error in setting the cache, return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Compose response message
	resp := u.Message(true, "Logged in")
	// Attach user to the response message
	resp["user"] = user

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
func GetUserByID(id uint) *User {

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
