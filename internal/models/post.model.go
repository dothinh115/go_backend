package models

import (
	"encoding/json"
	"project/internal/database"
	"time"

	"github.com/google/uuid"
)

func init() {
	database.GetDb().AutoMigrate(&PostCategory{}, &Post{})
}

type PostCategory struct {
	PostID     uint `gorm:"column:post_id;primaryKey"`
	CategoryID uint `gorm:"column:category_id;primaryKey"`
}

func (PostCategory) TableName() string {
	return "post_categories"
}

type Post struct {
	ID         uint       `gorm:"column:id;primaryKey" json:"id"`
	CreatedAt  time.Time  `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt  time.Time  `gorm:"column:updatedAt" json:"updatedAt"`
	Content    string     `gorm:"column:content;type:not null" json:"content"`
	AuthorID   uuid.UUID  `gorm:"column:authorId;type:uuid,not null" json:"-"`
	Author     User       `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Categories []Category `gorm:"many2many:post_categories;joinForeignKey:PostID;joinReferences:CategoryID;constraint:OnCreate:CASCADE" json:"categories"`
}

func (*Post) TableName() string {
	return "post"
}

type CreatePost struct {
	Author     uuid.UUID `json:"author"`
	Content    string    `json:"content"`
	Categories []uint    `json:"categories"`
}

func (p Post) MarshalJSON() ([]byte, error) {
	type Alias Post
	if p.Author.ID.String() == (uuid.UUID{}).String() {
		return json.Marshal(&struct {
			*Alias
			Author interface{} `json:"author,omitempty"`
		}{
			Alias:  (*Alias)(&p),
			Author: nil,
		})
	}
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(&p),
	})
}
