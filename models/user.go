package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key" example:"550e8400-e29b-41d4-a716-446655440000"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	Name      string         `json:"name" example:"John Doe"`
	Email     string         `json:"email" gorm:"uniqueIndex" example:"john@example.com"`
	Password  string         `json:"-" example:"password123"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null" example:"johndoe"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}
