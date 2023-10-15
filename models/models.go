package models

import "time"

type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserID      int       `json:"user_id"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateArticle struct {
	Title    string `json:"title"`
	Contents string `json:"contents"`
}

type User struct {
	UserID    int       `json:"user_id"`
	GoogleID  string    `json:"google_id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Nice struct {
	NiceID    int       `json:"nice_id"`
	UserID    int       `json:"user_id"`
	ArticleID int       `json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}
