package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
	"golang.org/x/oauth2"
)

func (s *MyAppService) GoogleCallbackService(code string) (*oauth2.Token, error) {

	token, err := s.config.Exchange(context.Background(), code)
	if err != nil {
		err = apperrors.ExchangeTokenFailed.Wrap(err, "fail to exchange code for token")
		return nil, err
	}

	client := s.config.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		err = apperrors.GetUserInfoFailed.Wrap(err, "fail to get user info")
		return nil, err
	}
	defer response.Body.Close()

	var userInfo models.GoogleUserDataResponse
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		err = apperrors.DecodeUserInfoFailed.Wrap(err, "fail to decode user info")
		return nil, err
	}

	err = repositories.ExistUser(s.db, userInfo)
	if errors.Is(err, sql.ErrNoRows) {
		_, err := repositories.InsertUser(s.db, userInfo)
		if err != nil {
			err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
			return nil, err
		}
		return token, nil
	} else if err != nil {
		return nil, err
	} else {
		return token, nil
	}
}

func (s *MyAppService) RegenerateAccessTokenService(refreshToken models.RefreshToken) (*oauth2.Token, error) {

	token := &oauth2.Token{
		RefreshToken: refreshToken.RefreshToken,
	}

	newToken, err := s.config.TokenSource(context.Background(), token).Token()
	if err != nil {
		fmt.Println("Failed to refresh token:", err)
		err = apperrors.RefreshTokenFailed.Wrap(err, "Failed to refresh token")
		return nil, err
	}

	return newToken, nil
}

func (s *MyAppService) UpdateUserService(userID int, updateUser models.UpdateUser) error {
	err := repositories.UpdateUser(s.db, userID, updateUser)

	if err != nil {
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update data")
		return err
	}

	return nil
}
