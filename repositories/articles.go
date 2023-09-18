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
	(?, ?, ?, now());
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
		select article_id, title, contents, username
		from articles
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
		rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserID)

		articleArray = append(articleArray, article)
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
	`

	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserID, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`
	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
