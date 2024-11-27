package models

import (
	"project/internal/database"
	"time"

	"github.com/google/uuid"
)

func init() {
	database.GetDb().AutoMigrate(&User{})
}

type User struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email     string    `gorm:"column:email;uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"column:password;not null" json:"-"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (User) TableName() string {
	return "user"
}
