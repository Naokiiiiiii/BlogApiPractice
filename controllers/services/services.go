package services

import (
	"github.com/Naokiiiiiii/BlogApiPractice/models"
)

type ArticleServicer interface {
	PostArticleService(article models.Article) (error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	UpdateArticleService(articleID int, updateArticle models.UpdateArticle) error
	DeleteArticleService(articleID int) error
}

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
	UpdateCommentService(commentID int, updateComment models.UpdateComment) error
	DeleteCommentService(commentID int) error
}

type NiceServicer interface {
	CreateOrDeleteNiceService(nice models.Nice) (models.Nice, error)
}

type UserServicer interface {
	GoogleCallbackService(code string) (models.GoogleOAuthToken, error)
	RegenerateAccessTokenService(refreshToken models.RefreshToken) (models.GoogleOAuthToken, error)
	GetUserService(accessToken string) (models.User, error)
	UpdateUserService(userID int, updateUser models.UpdateUser) error
}
