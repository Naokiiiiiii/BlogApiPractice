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

func SelectUser(db *sql.DB, email string) (models.User, error) {
	const sqlStr = `
		select * from users where email = ?;
	`

	row := db.QueryRow(sqlStr, email)
	if err := row.Err(); err != nil {
		return models.User{}, err
	}

	var user models.User
	var createdTime sql.NullTime
	var updatedTime sql.NullTime
	err := row.Scan(&user.UserID, &user.GoogleID, &user.UserName, &user.Email, &createdTime, &updatedTime)
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

func UpdateUser(db *sql.DB, userID int, updateUser models.UpdateUser) error {
	const sqlStr = `
		UPDATE users SET username = ?, updated_at = now() WHERE user_id = ?;
	`

	_, err := db.Exec(sqlStr, updateUser.UserName, userID)
	if err != nil {
		return err
	}

	return nil
}
