package models

import (
	"encoding/json"
	"time"
)

type Post struct {
	ID              *int             `json:"id" db:"id"`
	Title           *string          `json:"title" db:"title"`
	Slug            *string          `json:"slug" db:"slug"`
	Content         *string          `json:"content" db:"content"`
	Excerpt         *string          `json:"excerpt" db:"excerpt"`
	Image           *string          `json:"image" db:"image"`
	PublishedAt     *time.Time       `json:"published_at" db:"published_at"`
	MetaTitle       *string          `json:"meta_title" db:"meta_title"`
	MetaDescription *string          `json:"meta_description" db:"meta_description"`
	Category        *string          `json:"category" db:"category"`
	Tags            *json.RawMessage `json:"tags" db:"tags"`
	Categories      *json.RawMessage `json:"categories" db:"categories"`
}

type PostCategory struct {
	Title *string
}
