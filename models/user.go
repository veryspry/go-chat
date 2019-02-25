package models

import (
	"errors"
	"fmt"
	"os"

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

// Create a new user
func (user *User) Create() (map[string]interface{}, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	db := GetDB()
	db.Create(&user)

	if user.ID <= 0 {
		return nil, errors.New("failed to create user, connection error")
	}

	//Create new JWT token for the newly registered user
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	user.Password = "" //delete password

	response := u.Message(true, "user has been created")
	response["user"] = user
	return response, nil
}

// GetUserByEmail - return a single user by email address
func GetUserByEmail(email string) (*User, error) {

	db := GetDB()

	user := &User{}

	// Find the user
	db.Table("users").Where("email = ?", email).First(user)
	if user.Email == "" {
		return user, errors.New("user not found")
	}

	return user, nil
}

// Login a user
func Login(email, password string) map[string]interface{} {

	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}

	// Success - don't send back a password
	user.Password = ""

	// Create JWT token
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	// Store the token in the response
	user.Token = tokenString

	resp := u.Message(true, "Logged In")
	resp["account"] = user

	fmt.Println(resp["account"])
	return resp
}
