package model

import "time"

type Comment struct {
	Content string
	Author  string
}

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPublished Status = "published"
)

func (s Status) IsValid() bool {
	return s == StatusDraft || s == StatusPublished
}

// Article represents a blog article
type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tag       []string  `json:"tag"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comment   []Comment `json:"comment"`
	Like      int       `json:"like"`
	Hot       int       `json:"hot"`
	UserID    int64     `json:"user_id"`
	Author    string    `json:"author"`
}
