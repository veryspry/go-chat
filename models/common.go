package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// BaseFields is a utility to extend other models
type BaseFields struct {
	ID uuid.UUID `gorm:"primary_key;not null" json:"id"`
	// `gorm:"primary_key;unique;not null"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}
