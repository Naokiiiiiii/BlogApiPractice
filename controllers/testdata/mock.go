package testdata

import "github.com/Naokiiiiiii/BlogApiPractice/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) PostArticleService(article models.Article) (models.Article, error) {
	return articleTestData[1], nil
}

func (s *serviceMock) GetArticleListService(page int) ([]models.Article, error) {
	return articleTestData, nil
}

func (s *serviceMock) GetArticleService(articleID int) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) UpdateArticleService(articleID int, article models.Article) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) DeleteArticleService(articleID int) error {
	return nil
}

func (s *serviceMock) PostNiceService(article models.Article) (models.Article, error) {
	return articleTestData[0], nil
}

func (s *serviceMock) PostCommentService(comment models.Comment) (models.Comment, error) {
	return commentTestData[0], nil
}

func (s *serviceMock) UpdateCommentService(commentID int, comment models.Comment) (models.Comment, error) {
	return commentTestData[0], nil
}

func (s *serviceMock) DeleteCommentService(commentID int) error {
	return nil
}
