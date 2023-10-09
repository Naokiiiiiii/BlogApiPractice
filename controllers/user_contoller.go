package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Naokiiiiiii/BlogApiPractice/controllers/services"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type UserController struct {
	service services.UserServicer
}

func NewUserController(s services.UserServicer) *UserController {
	return &UserController{service: s}
}

func (c *UserController) GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIANT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIANT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Endpoint:     google.Endpoint,
		Scopes:       []string{"profile", "email"},
	}
	url := config.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

func (c *UserController) GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIANT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIANT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Endpoint:     google.Endpoint,
		Scopes:       []string{"profile", "email"},
	}

	code := r.URL.Query().Get("code")
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange code for token", http.StatusInternalServerError)
		return
	}

	client := config.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
		return
	}

	fmt.Println("userinfo", userInfo["id"], userInfo["email"], userInfo["name"])

	// DBにユーザー情報格納

	// アクセストークン、リフレッシュトークンを返す
}
