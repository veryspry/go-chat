package models

import (
	"strings"
)

// NormalizeUsers scans the user table and ensures that email, firstname and lastname are all stored as lowercase values
func NormalizeUsers() {
	db := GetDB()
	users := []*User{}
	db.Find(&users)
	for _, u := range users {
		u.Email = strings.ToLower(u.Email)
		u.FirstName = strings.ToLower(u.FirstName)
		u.FirstName = strings.ToLower(u.FirstName)
		u.LastName = strings.ToLower(u.LastName)
		db.Save(&u)
	}
}
