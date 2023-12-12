package models

import "time"

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserID      int       `json:"user_id"`
	UserName    string    `json:"user_name"`
	NiceList    []Nice    `json:"nice"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateArticle struct {
	Title    string `json:"title"`
	Contents string `json:"contents"`
}
