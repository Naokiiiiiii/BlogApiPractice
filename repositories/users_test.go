package repositories_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories/testdata"
)

func TestInsertUser(t *testing.T) {
	user := models.GoogleUserDataResponse{
		Id:    "123123123",
		Name:  "test",
		Email: "test@test.com",
	}

	expectedUserName := "test"
	newUser, err := repositories.InsertUser(testDB, user)
	if err != nil {
		t.Error(err)
	}

	if newUser.UserName != expectedUserName {
		t.Errorf("new user name is expected %s but got %s\n", expectedUserName, newUser.UserName)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from users
			where  google_id = ? and email = ? and username = ?
		`
		testDB.Exec(sqlStr, user.Id, user.Email, user.Name)
	})

}

func TestSelectUser(t *testing.T) {
	expectedUserData := testdata.UserTestData

	got, err := repositories.SelectUser(testDB, expectedUserData.Email)
	if err != nil {
		t.Fatal(err)
	}

	if expectedUserData.GoogleID != got.GoogleID {
		t.Errorf("GoogleID: get %s but want %s\n", got.GoogleID, expectedUserData.GoogleID)
	}

	if expectedUserData.UserName != got.UserName {
		t.Errorf("UserName: get %s but want %s\n", got.UserName, expectedUserData.UserName)
	}

	if expectedUserData.Email != got.Email {
		t.Errorf("UserName: get %s but want %s\n", got.Email, expectedUserData.Email)
	}
}

func TestExistUser(t *testing.T) {
	existEmail := models.GoogleUserDataResponse{
		Email: "exsample@gmail.com",
	}

	err := repositories.ExistUser(testDB, existEmail)

	if errors.Is(err, sql.ErrNoRows) {
		t.Errorf("email address %s that should not exist, but it does.", existEmail.Email)
	}

	noExistEmail := models.GoogleUserDataResponse{
		Email: "noexist@gmail.com",
	}

	err = repositories.ExistUser(testDB, noExistEmail)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("email address %s that should not exist, but it does.", noExistEmail.Email)
	}
}

func TestUpdateUser(t *testing.T) {
	userID := 1
	expectedUserData := testdata.UpdateUserTestData

	err := repositories.UpdateUser(testDB, userID, expectedUserData)

	if err != nil {
		t.Error(err)
	}
}
