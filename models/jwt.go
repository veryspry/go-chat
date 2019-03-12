package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Token JWT claims struct
type Token struct {
	UserID uuid.UUID
	jwt.StandardClaims
}
