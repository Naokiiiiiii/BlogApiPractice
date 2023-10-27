package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Naokiiiiiii/BlogApiPractice/apperrors"
	"github.com/Naokiiiiiii/BlogApiPractice/controllers/services"
	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

type UserController struct {
	service services.UserServicer
	config  oauth2.Config
}

func NewUserController(s services.UserServicer, config oauth2.Config) *UserController {
	return &UserController{service: s, config: config}
}

func (c *UserController) GoogleLoginHandler(w http.ResponseWriter, req *http.Request) {
	url := c.config.AuthCodeURL("", oauth2.AccessTypeOffline)
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

func (c *UserController) SelectUserInfoHandler(w http.ResponseWriter, req *http.Request) {
	authorizationHeader := req.Header.Get("Authorization")

	authHeaders := strings.Split(authorizationHeader, " ")

	if len(authHeaders) == 2 && authHeaders[0] == "Bearer" {
		accessToken := authHeaders[1]
		user, _ := c.service.GetUserService(accessToken)
		json.NewEncoder(w).Encode(user)
	}

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
