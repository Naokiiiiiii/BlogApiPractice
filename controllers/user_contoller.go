package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/controllers/services"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/gorilla/mux"
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
		RedirectURL:  "http://localhost:5173",
		Endpoint:     google.Endpoint,
		Scopes:       []string{"profile", "email"},
	}
	url := config.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, req, url, http.StatusFound)
}

func (c *UserController) GoogleTokenHandler(w http.ResponseWriter, req *http.Request) {

	var googleOAuthCode models.GoogleOAuthCode
	if err := json.NewDecoder(req.Body).Decode(&googleOAuthCode); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	token, err := c.service.GoogleCallbackService(googleOAuthCode.Code)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

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

func (c *UserController) UpdateUserHandler(w http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "queryparam must be number")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	var reqUpdateUser models.UpdateUser
	if err := json.NewDecoder(req.Body).Decode(&reqUpdateUser); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad resuest body")
		apperrors.ErrorHandler(w, req, err)
	}

	err = c.service.UpdateUserService(userID, reqUpdateUser)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(err)
}
