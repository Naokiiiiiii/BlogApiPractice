package repositories_test

import (
	"testing"
	"time"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticle(t *testing.T) {
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
				t.Errorf("UserID: get %d but want %d\n", got.UserID, test.expected.UserID)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
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
		t.Errorf("new article title is expected %s but got %s\n", expectedArticleTitle, newArticle.Title)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and user_id = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserID)
	})
}

func TestUpdateArticle(t *testing.T) {
	article := models.Article{
		ID:        1,
		Title:     "Updated Title",
		Contents:  "Updated Contents",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	resultArticle, err := repositories.UpdateArticle(testDB, article, article.ID)
	if err != nil {
		t.Error(err)
	}

	if article.Title != resultArticle.Title {
		t.Errorf("new article title is expected %s but got %s\n", article.Title, resultArticle.Title)
	}
}

func TestDeleteArticle(t *testing.T) {
	deleteArticleID := 3

	err := repositories.DeleteArticle(testDB, deleteArticleID)

	if err != nil {
		t.Errorf("DeleteArticle returned an error: %v", err)
	}
}
