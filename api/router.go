package api

import (
	"database/sql"
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
	nCon := controllers.NewNiceController(ser)
	uCon := controllers.NewUserController(ser)
	r := mux.NewRouter()

	// 認証が必要ないAPI
	r.HandleFunc("/login", uCon.GoogleLoginHandler)
	r.HandleFunc("/callback", uCon.GoogleCallbackHandler)
	r.HandleFunc("/regenerateToken", uCon.RegenerateAccessTokenHandler)

	// 認証が必要なAPI
	authRequired := r.PathPrefix("/").Subrouter()
	authRequired.Use(middlewares.AuthMiddleware)

	authRequired.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	authRequired.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	authRequired.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	authRequired.HandleFunc("/article/nice", nCon.PostNiceHandler).Methods(http.MethodPost)
	authRequired.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}
