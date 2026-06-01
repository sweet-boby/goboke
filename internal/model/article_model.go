package model

import "time"

type Comment struct {
	Content string
	Author  string
}

type ArticleStatus string

const (
	StatusDraft     ArticleStatus = "draft"
	StatusPublished ArticleStatus = "published"
	StatusDeleted   ArticleStatus = "deleted"
)

func (s ArticleStatus) IsValid() bool {
	return s == StatusDraft || s == StatusPublished
}

// Article represents a blog article
type Article struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	Tag       []string      `json:"tag"`
	Status    ArticleStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Comment   []Comment     `json:"comment"`
	Like      int           `json:"like"`
	Hot       int           `json:"hot"`
	UserID    int           `json:"user_id"`
	Author    string        `json:"author"`
}
