package dto

import "github.com/google/uuid"

type CreatePost struct {
	Author     uuid.UUID `json:"author"`
	Content    string    `json:"content"`
	Categories []uint    `json:"categories"`
}
