package controllers_test

import (
	"testing"

	"github.com/Naokiiiiiii/BlogApiPractice/controllers"
	"github.com/Naokiiiiiii/BlogApiPractice/controllers/testdata"

	_ "github.com/go-sql-driver/mysql"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {

	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}
