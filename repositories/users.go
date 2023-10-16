package repositories

import (
	"database/sql"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
)

func InsertUser(db *sql.DB, googleUser models.GoogleUserDataResponse) (models.User, error) {

	const sqlStr = `insert into users (google_id, username, email, created_at) values
		(?, ?, ?, now());
	`

	var newUser models.User
	newUser.GoogleID, newUser.UserName, newUser.Email = googleUser.Id, googleUser.Name, googleUser.Email

	result, err := db.Exec(sqlStr, googleUser.Id, googleUser.Name, googleUser.Email)
	if err != nil {
		return models.User{}, err
	}

	id, _ := result.LastInsertId()
	newUser.UserID = int(id)

	return newUser, nil
}

func GetUser(db *sql.DB, userID int) (models.User, error) {
	const sqlStr = `
		select * from users where email = ?;	
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

func ExistUser(db *sql.DB, googleUser models.GoogleUserDataResponse) error {
	const sqlStr = `
		select * from users where email = ?;
	`
	var user models.User
	row := db.QueryRow(sqlStr, googleUser.Email)
	err := row.Scan(&user.Email, &user.UserID, &user.UserName, &user.GoogleID, &user.CreatedAt, &user.UpdatedAt)

	return err
}
