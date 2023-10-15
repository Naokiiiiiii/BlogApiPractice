package services

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
)

func (s *MyAppService) CreateOrDeleteNiceService(nice models.Nice) (models.Nice, error) {

	err := repositories.ExistNice(s.db, nice)
	fmt.Println(err)
	if errors.Is(err, sql.ErrNoRows) {
		newNice, err := repositories.InsertNice(s.db, nice)
		if err != nil {
			err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
			return models.Nice{}, err
		}
		return newNice, nil
	} else {
		err := repositories.DeleteNice(s.db, nice)
		if err != nil {
			err = apperrors.DeleteDataFailed.Wrap(err, "fail to delete data")
			return models.Nice{}, err
		}
		return models.Nice{}, nil
	}
}
