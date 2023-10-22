package repositories

import (
	"database/sql"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
)

func InsertUser(db *sql.DB, googleUser models.GoogleUserDataResponse) (models.User, error) {

	const sqlStr = `insert into users (google_id, username, email, created_at, updated_at) values
		(?, ?, ?, now(), now());
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

func GetUser(db *sql.DB, googleID int) (models.User, error) {
	const sqlStr = `
		select * from users where google_id = ?;
	`

	row := db.QueryRow(sqlStr, googleID)
	if err := row.Err(); err != nil {
		return models.User{}, err
	}

	var user models.User
	var createdTime sql.NullTime
	var updatedTime sql.NullTime
	err := row.Scan(&user.UserID, &user.UserName, &user.GoogleID, &createdTime, &updatedTime)
	if err != nil {
		return models.User{}, err
	}

	if createdTime.Valid {
		user.CreatedAt = createdTime.Time
	}

	if updatedTime.Valid {
		user.UpdatedAt = updatedTime.Time
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
