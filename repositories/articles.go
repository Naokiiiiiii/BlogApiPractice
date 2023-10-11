package repositories

import (
	"database/sql"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
)

const (
	articleNumPerPage = 5
)

func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
	insert into articles (title, contents, user_id, created_at) values
	(?, ?, ?, now(), now());
	`

	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserID = article.Title, article.Contents, article.UserID

	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserID)
	if err != nil {
		return models.Article{}, err
	}

	id, _ := result.LastInsertId()

	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select articles.article_id, articles.title, articles.contents, articles.user_id, users.username
		from articles
		inner join users on articles.user_id = users.user_id
		limit ? offset ?;
	`

	rows, err := db.Query(sqlStr, articleNumPerPage, ((page - 1) * articleNumPerPage))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserID, &article.UserName)

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select articles.*, users.username
		from articles
		inner join users on articles.user_id = users.user_id
		where article_id = ?;
	`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserID, &createdTime, &article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}
