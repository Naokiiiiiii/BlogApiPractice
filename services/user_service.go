package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func (s *MyAppService) GoogleCallbackService(code string) (*oauth2.Token, error) {

	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIANT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIANT_SECRET"),
		RedirectURL:  "http://localhost:5173",
		Endpoint:     google.Endpoint,
		Scopes:       []string{"profile", "email"},
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		err = apperrors.ExchangeTokenFailed.Wrap(err, "fail to exchange code for token")
		return nil, err
	}

	client := config.Client(context.Background(), token)
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
	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIANT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIANT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Endpoint:     google.Endpoint,
		Scopes:       []string{"profile", "email", "offline"},
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken.RefreshToken,
	}

	newToken, err := config.TokenSource(context.Background(), token).Token()
	if err != nil {
		fmt.Println("Failed to refresh token:", err)
		err = apperrors.RefreshTokenFailed.Wrap(err, "Failed to refresh token")
		return nil, err
	}

	return newToken, nil
}
