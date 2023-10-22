package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/controllers/services"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type UserController struct {
	service services.UserServicer
}

func NewUserController(s services.UserServicer) *UserController {
	return &UserController{service: s}
}

func (c *UserController) GoogleLoginHandler(w http.ResponseWriter, req *http.Request) {
	config := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIANT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIANT_SECRET"),
		RedirectURL:  "http://localhost:8080/callback",
		Endpoint:     google.Endpoint,
		Scopes:       []string{"profile", "email"},
	}
	url := config.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, req, url, http.StatusFound)
}

func (c *UserController) GoogleCallbackHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("reqURL", req.URL)
	code := req.URL.Query().Get("code")
	token, err := c.service.GoogleCallbackService(code)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	fmt.Println("token", token.RefreshToken)

	json.NewEncoder(w).Encode(token)
}

func (c *UserController) RegenerateAccessTokenHandler(w http.ResponseWriter, req *http.Request) {

	var refreshToken models.RefreshToken
	if err := json.NewDecoder(req.Body).Decode(&refreshToken); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	newToken, err := c.service.RegenerateAccessTokenService(refreshToken)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(newToken)
}
