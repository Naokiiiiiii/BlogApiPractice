package repositories_test

import (
	"testing"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories/testdata"
)

func TestSelectNiceList(t *testing.T) {
	expectedNum := len(testdata.NiceTestData)
	got, err := repositories.SelectArticleNiceList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d nices \n", expectedNum, num)
	}
}

func TestInsertNice(t *testing.T) {
	nice := models.Nice{
		UserID:    1,
		ArticleID: 2,
	}

	expectedNiceID := 2
	newNice, err := repositories.InsertNice(testDB, nice)
	if err != nil {
		t.Error(err)
	}
	if newNice.ArticleID == expectedNiceID {
		t.Errorf("new article id is expected %d but got %d", expectedNiceID, newNice.NiceID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from nices
			where user_id = ? and article_id = ?
		`
		testDB.Exec(sqlStr, nice.UserID, nice.ArticleID)
	})
}
