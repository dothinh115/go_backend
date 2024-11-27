package models

import (
	"project/internal/database"
	"time"
)

func init() {
	database.GetDb().AutoMigrate(&Category{})
}

type Category struct {
	ID        uint      `gorm:"column:id;primaryKLey" json:"id"`
	Title     string    `gorm:"column:title;not null" json:"title"`
	Slug      string    `gorm:"column:slug;not null" json:"slug"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (Category) TableName() string {
	return "category"
}
