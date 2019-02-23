package models

import (
	"os"

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
func (user *User) Create() map[string]interface{} {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB.Create(&user)

	if user.ID <= 0 {
		return u.Message(false, "Failed to create user, connection error.")
	}

	//Create new JWT token for the newly registered user
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	user.Password = "" //delete password

	response := u.Message(true, "user has been created")
	response["user"] = user
	return response
}
