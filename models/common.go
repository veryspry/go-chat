package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// BaseFields is a utility to extend other models
type BaseFields struct {
	ID uuid.UUID `gorm:"primary_key;not null"`
	// `gorm:"primary_key;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
