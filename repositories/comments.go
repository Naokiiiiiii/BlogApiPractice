package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, user_id, message, created_at, updated_at) values
		(?, ?, ?, now(), now());
	`
	var newComment models.Comment
	newComment.ArticleID, newComment.UserID, newComment.Message = comment.ArticleID, comment.UserID, comment.Message

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.UserID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, _ := result.LastInsertId()
	newComment.CommentID = int(id)

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select comments.*, users.username
		from comments
		inner join users on comments.user_id = users.user_id
		where article_id = ?;
	`

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		var updatedTime sql.NullTime
		rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.UserID, &comment.UserName, &comment.Message, &createdTime, &updatedTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		if updatedTime.Valid {
			comment.UpdatedAt = updatedTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}

func UpdateComment(db *sql.DB, updateComment models.UpdateComment, commentID int) error {
	const sqlStr = `
		UPDATE comments SET message = ?, updated_at = now() WHERE comment_id = ?;
	`

	_, err := db.Exec(sqlStr, updateComment.Message, commentID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func DeleteComment(db *sql.DB, commnetID int) error {
	const sqlStr = `DELETE FROM comments WHERE comment_id = ?`

	_, err := db.Exec(sqlStr, commnetID)

	if err != nil {
		return err
	}

	return nil
}
