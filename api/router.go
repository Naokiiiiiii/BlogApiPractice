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

	r.Use(middlewares.CorsMiddleware)

	// 認証が必要ないAPI
	r.HandleFunc("/login", uCon.GoogleLoginHandler)
	r.HandleFunc("/token", uCon.GoogleTokenHandler)
	r.HandleFunc("/regenerateToken", uCon.RegenerateAccessTokenHandler).Methods(http.MethodPost)

	r.HandleFunc("/article/{id:[0-9]+}", aCon.UpdateArticleHandler).Methods(http.MethodPut)
	r.HandleFunc("/comment/{id:[0-9]+}", cCon.UpdateCommentHandler).Methods(http.MethodPut)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.DeleteArticleHandler).Methods(http.MethodDelete)
	r.HandleFunc("/comment/{id:[0-9]+}", cCon.DeleteCommentHandler).Methods(http.MethodDelete)
	r.HandleFunc("/article/nice", nCon.CreateOrDeleteNiceHandler).Methods(http.MethodPost)

	// 認証が必要なAPI
	authRequired := r.PathPrefix("/").Subrouter()
	authRequired.Use(middlewares.AuthMiddleware)

	authRequired.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	authRequired.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	authRequired.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	authRequired.HandleFunc("/article/nice", nCon.CreateOrDeleteNiceHandler).Methods(http.MethodPost)
	authRequired.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	return r
}
