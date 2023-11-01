package repositories_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
)

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
