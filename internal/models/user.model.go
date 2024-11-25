package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"column:CreatedAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt" json:"updatedAt"`
}

func (User) TableName() string {
	return "user"
}
