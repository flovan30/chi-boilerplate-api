package entity

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `gorm:"type:uuid"`
	Name      string    `gorm:"uniqueIndex;not null"`
	Resume    string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
