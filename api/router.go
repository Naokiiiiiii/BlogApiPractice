package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Naokiiiiiii/BlogApiPractice/api/middlewares"
	"github.com/Naokiiiiiii/BlogApiPractice/controllers"
	"github.com/Naokiiiiiii/BlogApiPractice/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)
	r := mux.NewRouter()

	// 認証が必要ないAPI
	r.HandleFunc("/callback", handleGoogleCallback)

	// 認証が必要なAPI
	authRequired := r.PathPrefix("/").Subrouter()
	authRequired.Use(middlewares.AuthMiddleware)

	authRequired.HandleFunc("/hello", aCon.HelloWorldHandler).Methods(http.MethodGet)
	authRequired.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	authRequired.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	authRequired.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	authRequired.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	authRequired.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	fmt.Println("redirect", code)
}
