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
	r := mux.NewRouter()

	r.HandleFunc("/hello", aCon.HelloWorldHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.AuthMiddleware)

	return r
}
