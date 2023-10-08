package repositories

import (
	"database/sql"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
)

func InsertUser(db *sql.DB, user models.User) (models.User, error) {

	const sqlStr = `insert into users (google_id, username, email, created_at) values
		(?, ?, ?, now());
	`

	var newUser models.User
	newUser.GoogleID, newUser.UserName, newUser.Email = user.GoogleID, user.UserName, user.Email

	result, err := db.Exec(sqlStr, user.GoogleID, user.UserName, user.Email)
	if err != nil {
		return models.User{}, err
	}

	id, _ := result.LastInsertId()
	newUser.UserID = int(id)

	return newUser, nil
}

func GetUser(db *sql.DB, userID int) (models.User, error) {
	const sqlStr = `
		select * from users where user_id = ?;	
	`

	row := db.QueryRow(sqlStr, userID)
	if err := row.Err(); err != nil {
		return models.User{}, err
	}

	var user models.User
	var createdTime sql.NullTime
	err := row.Scan(&user.UserID, &user.UserName, &user.GoogleID, &createdTime)
	if err != nil {
		return models.User{}, err
	}

	if createdTime.Valid {
		user.CreatedAt = createdTime.Time
	}

	return user, nil
}
