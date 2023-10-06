package repositories_test

import (
	"testing"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
)

func TestSelectNiceList(t *testing.T) {
	articleID := 1
	got, err := repositories.SelectArticleNiceList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, nice := range got {
		if nice.ArticleID != articleID {
			t.Errorf("want nice of articleID %d but got ID %d\n", articleID, nice.ArticleID)
		}
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
