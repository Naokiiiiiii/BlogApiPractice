package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Naokiiiiiii/BlogApiPractice/api/middlewares"
	"github.com/Naokiiiiiii/BlogApiPractice/controllers"
	"github.com/Naokiiiiiii/BlogApiPractice/services"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)
	nCon := controllers.NewNiceController(ser)
	r := mux.NewRouter()

	// 認証が必要ないAPI
	r.HandleFunc("/login", handleGoogleLogin)
	r.HandleFunc("/callback", handleGoogleCallback)

	// 認証が必要なAPI
	authRequired := r.PathPrefix("/").Subrouter()
	authRequired.Use(middlewares.AuthMiddleware)

	authRequired.HandleFunc("/hello", aCon.HelloWorldHandler).Methods(http.MethodGet)
	authRequired.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	authRequired.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	authRequired.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	authRequired.HandleFunc("/article/nice", nCon.PostNiceHandler).Methods(http.MethodPost)
	authRequired.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
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

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println("userinfo", userInfo)
}
