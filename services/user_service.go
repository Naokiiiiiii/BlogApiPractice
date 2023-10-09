package services

import (
	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
)

func (s *MyAppService) PostUserService(user models.User) (models.User, error) {
	newUser, err := repositories.InsertUser(s.db, user)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.User{}, err
	}

	return newUser, nil
}
