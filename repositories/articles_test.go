package repositories_test

import (
	"testing"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectarticle(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		},
		{
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserID != test.expected.UserID {
				t.Errorf("UserName: get %d but want %d\n", got.UserID, test.expected.UserID)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserID:   1,
	}

	expectedArticleTitle := "insertTest"
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}
	if newArticle.Title != expectedArticleTitle {
		t.Errorf("new article id is expected %s but got %s\n", expectedArticleTitle, newArticle.Title)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and user_id = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserID)
	})
}

func TestUpdateNiceNum(t *testing.T) {
	articleID := 1
	before, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	err = repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	after, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	if after.NiceNum-before.NiceNum != 1 {
		t.Error("fail to update nice num")
	}

}
