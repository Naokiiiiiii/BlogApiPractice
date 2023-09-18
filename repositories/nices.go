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
