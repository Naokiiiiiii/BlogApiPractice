package repositories

import (
	"database/sql"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
)

func SelectArticleNiceList(db *sql.DB, articleID int) ([]models.Nice, error) {
	const sqlStr = `
		select *
		from nices
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	NiceArray := make([]models.Nice, 0)
	for rows.Next() {
		var nice models.Nice
		var createdTime sql.NullTime
		rows.Scan(&nice.NiceID, &nice.ArticleID, &nice.UserID, &createdTime)

		if createdTime.Valid {
			nice.CreatedAt = createdTime.Time
		}

		NiceArray = append(NiceArray, nice)
	}

	return NiceArray, nil
}

func ExistNice(db *sql.DB, nice models.Nice) error {
	const sqlStr = `
		select * from nices where article_id = ? and user_id = ?;
	`
	var result models.Nice
	row := db.QueryRow(sqlStr, nice.ArticleID, nice.UserID)
	err := row.Scan(&result.NiceID, &result.UserID, &result.ArticleID, &result.CreatedAt)

	return err
}

func InsertNice(db *sql.DB, nice models.Nice) (models.Nice, error) {
	const sqlStr = `insert into nices (article_id, user_id, created_at) values
		(?, ?, now());
	`

	var newNice models.Nice
	newNice.UserID, newNice.ArticleID = nice.ArticleID, nice.UserID

	result, err := db.Exec(sqlStr, nice.ArticleID, nice.UserID)
	if err != nil {
		return models.Nice{}, err
	}

	id, _ := result.LastInsertId()
	newNice.NiceID = int(id)

	return newNice, nil
}

func DeleteNice(db *sql.DB, nice models.Nice) error {
	const sqlStr = `DELETE FROM nices where article_id = ? and user_id = ?;`
	_, err := db.Exec(sqlStr, nice.ArticleID, nice.UserID)

	if err != nil {
		return err
	}

	return nil
}
