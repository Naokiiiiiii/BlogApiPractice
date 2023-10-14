package services

import (
	"database/sql"
	"errors"

	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
)

func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var niceList []models.Nice
	var articleGetErr, commentGetErr, niceGetErr error

	type articleResult struct {
		article models.Article
		err     error
	}

	articleChan := make(chan articleResult)
	defer close(articleChan)

	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		article, err := repositories.SelectArticleDetail(db, articleID)
		ch <- articleResult{article: article, err: err}
	}(articleChan, s.db, articleID)

	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}

	commentChan := make(chan commentResult)
	defer close(commentChan)

	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		commentList, err := repositories.SelectCommentList(db, articleID)
		ch <- commentResult{commentList: &commentList, err: err}
	}(commentChan, s.db, articleID)

	type niceResult struct {
		niceList *[]models.Nice
		err      error
	}

	niceChan := make(chan niceResult)
	defer close(niceChan)
	go func(ch chan<- niceResult, db *sql.DB, articleID int) {
		niceList, err := repositories.SelectArticleNiceList(db, articleID)
		ch <- niceResult{niceList: &niceList, err: err}
	}(niceChan, s.db, articleID)

	for i := 0; i < 3; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = *cr.commentList, cr.err
		case nr := <-niceChan:
			niceList, niceGetErr = *nr.niceList, nr.err
		}
	}

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "no data")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "fail to get data")
		return models.Article{}, err
	}

	if niceGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(niceGetErr, "fail to get data")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)
	article.NiceNum = len(niceList)

	return article, nil
}

func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {

	newArticle, err := repositories.InsertArticle(s.db, article)

	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {

	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "nodata")
		return nil, err
	}

	return articleList, nil
}

func (s *MyAppService) UpdateArticleService(articleID int, article models.Article) (models.Article, error) {
	newArticle, err := repositories.UpdateArticle(s.db, article, articleID)
	if err != nil {
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update data")
		return models.Article{}, err
	}

	return newArticle, nil
}

func (s *MyAppService) DeleteArticleService(articleID int) error {
	err := repositories.DeleteArticle(s.db, articleID)
	if err != nil {
		err = apperrors.DeleteDataFailed.Wrap(err, "fail to delete data")
		return err
	}

	return nil
}
