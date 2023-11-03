package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

func (s *MyAppService) GoogleCallbackService(code string) (models.GoogleOAuthToken, error) {

	token, err := s.config.Exchange(context.Background(), code)
	if err != nil {
		err = apperrors.ExchangeTokenFailed.Wrap(err, "fail to exchange code for token")
		return models.GoogleOAuthToken{}, err
	}

	client := s.config.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		err = apperrors.GetUserInfoFailed.Wrap(err, "fail to get user info")
		return models.GoogleOAuthToken{}, err
	}
	defer response.Body.Close()

	var userInfo models.GoogleUserDataResponse
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		err = apperrors.DecodeUserInfoFailed.Wrap(err, "fail to decode user info")
		return models.GoogleOAuthToken{}, err
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		err = errors.New("")
		err = apperrors.GetIDTokenFailed.Wrap(err, "Failed to extract id_token")
		return models.GoogleOAuthToken{}, err
	}

	googleToken := models.GoogleOAuthToken{
		AccessToken:  token.AccessToken,
		IDToken:      idToken,
		RefreshToken: token.RefreshToken,
	}

	err = repositories.ExistUser(s.db, userInfo)
	if errors.Is(err, sql.ErrNoRows) {
		_, err := repositories.InsertUser(s.db, userInfo)
		if err != nil {
			err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
			return models.GoogleOAuthToken{}, err
		}
		return googleToken, nil
	} else if err != nil {
		return models.GoogleOAuthToken{}, err
	} else {
		return googleToken, nil
	}
}

func (s *MyAppService) RegenerateAccessTokenService(refreshToken models.RefreshToken) (models.GoogleOAuthToken, error) {

	token := &oauth2.Token{
		RefreshToken: refreshToken.RefreshToken,
	}

	newToken, err := s.config.TokenSource(context.Background(), token).Token()
	if err != nil {
		err = apperrors.ExchangeRefreshTokenFailed.Wrap(err, "Failed to refresh token")
		return models.GoogleOAuthToken{}, err
	}

	idToken, ok := newToken.Extra("id_token").(string)
	if !ok {
		err = errors.New("")
		err = apperrors.GetIDTokenFailed.Wrap(err, "Failed to extract id_token")
		return models.GoogleOAuthToken{}, err
	}

	googleToken := models.GoogleOAuthToken{
		AccessToken:  newToken.AccessToken,
		IDToken:      idToken,
		RefreshToken: newToken.RefreshToken,
	}

	return googleToken, nil
}

func (s *MyAppService) GetUserService(idToken string) (models.User, error) {

	payload, err := idtoken.ParsePayload(idToken)

	if err != nil {
		err = apperrors.ParsePayloadFailed.Wrap(err, "fail to parse payload")
		return models.User{}, err
	}

	email := payload.Claims["email"].(string)

	user, err := repositories.SelectUser(s.db, email)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to update data")
		return models.User{}, err
	}

	return user, nil
}

func (s *MyAppService) UpdateUserService(userID int, updateUser models.UpdateUser) error {
	err := repositories.UpdateUser(s.db, userID, updateUser)

	if err != nil {
		err = apperrors.UpdateDataFailed.Wrap(err, "fail to update data")
		return err
	}

	return nil
}
