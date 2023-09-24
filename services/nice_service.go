package services

import (
	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
)

func (s *MyAppService) PostNiceService(nice models.Nice) (models.Nice, error) {
	newNice, err := repositories.InsertNice(s.db, nice)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Nice{}, err
	}

	return newNice, nil
}
