package services

import (
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"golang.org/x/oauth2"
)

type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	UpdateArticleService(articleID int, article models.Article) (models.Article, error)
	DeleteArticleService(articleID int) error
}

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
	UpdateCommentService(commentID int, comment models.Comment) (models.Comment, error)
	DeleteCommentService(commentID int) error
}

type NiceServicer interface {
	CreateOrDeleteNiceService(nice models.Nice) (models.Nice, error)
}

type UserServicer interface {
	PostUserService(user models.User) (models.User, error)
	GoogleCallbackService(code string) (*oauth2.Token, map[string]interface{}, error)
	RegenerateAccessTokenService(refreshToken models.RefreshToken) (*oauth2.Token, error)
}
