package model

import "time"

type ArticleHistory struct {
	ID        int       `json:"id"`
	LookAt    time.Time `json:"look_at"`
	UserID    int       `json:"user_id"`
	ArticleID int       `json:"article_id"`
}
