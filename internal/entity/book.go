package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID        uuid.UUID `gorm:"type:uuid"`
	Name      string    `gorm:"uniqueIndex;not null"`
	Resume    string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Book) BeforeCreate(_ *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}
