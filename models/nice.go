package models

import "time"

type Nice struct {
	NiceID    int       `json:"nice_id"`
	UserID    int       `json:"user_id"`
	ArticleID int       `json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
}
