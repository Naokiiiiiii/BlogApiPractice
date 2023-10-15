package services

import (
	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {

	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}

func (s *MyAppService) UpdateCommentService(commentID int, updateComment models.UpdateComment) error {
	err := repositories.UpdateComment(s.db, updateComment, commentID)
	if err != nil {
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update data")
		return err
	}

	return nil
}

func (s *MyAppService) DeleteCommentService(commentID int) error {
	err := repositories.DeleteComment(s.db, commentID)
	if err != nil {
		err = apperrors.DeleteDataFailed.Wrap(err, "fail to delete data")
		return err
	}

	return nil
}
