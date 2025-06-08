package models

import (
	"time"
)

// Post represents a blog post in the system.
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewPost creates a new Post instance with the current time as CreatedAt and UpdatedAt.
func NewPost(title, content string, authorID int) *Post {
	return &Post{
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
