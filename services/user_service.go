package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func (s *MyAppService) PostUserService(user models.User) (models.User, error) {
	newUser, err := repositories.InsertUser(s.db, user)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.User{}, err
	}

	return newUser, nil
}

func (s *MyAppService) GoogleCallbackService(code string) (*oauth2.Token, map[string]interface{}, error) {

	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIANT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIANT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Endpoint:     google.Endpoint,
		Scopes:       []string{"profile", "email"},
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		err = apperrors.ExchangeTokenFailed.Wrap(err, "fail to exchange code for token")
		return nil, nil, err
	}

	client := config.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		err = apperrors.GetUserInfoFailed.Wrap(err, "fail to get user info")
		return nil, nil, err
	}
	defer response.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		err = apperrors.DecodeUserInfoFailed.Wrap(err, "fail to decode user info")
		return nil, nil, err
	}

	fmt.Println("userinfo", userInfo["id"], userInfo["email"], userInfo["name"])

	return token, userInfo, nil
}

func (s *MyAppService) RegenerateAccessTokenService(refreshToken string) (*oauth2.Token, error) {
	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIANT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIANT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Endpoint:     google.Endpoint,
		Scopes:       []string{"profile", "email"},
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	newToken, err := config.TokenSource(context.Background(), token).Token()
	if err != nil {
		fmt.Println("Failed to refresh token:", err)
		return nil, err
	}

	return newToken, nil
}
