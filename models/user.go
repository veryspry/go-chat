package models

import "github.com/jinzhu/gorm"

// User type for db
type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
