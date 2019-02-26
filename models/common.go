package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// BaseModel - use to extend all other DB types
type BaseModel struct {
	ID        uuid.UUID `gorm:"primary_key;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
