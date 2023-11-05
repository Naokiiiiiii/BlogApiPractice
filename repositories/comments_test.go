package repositories_test

import (
	"testing"

	"github.com/Naokiiiiiii/BlogApiPractice/repositories"
	"github.com/Naokiiiiiii/BlogApiPractice/repositories/testdata"
)

func TestSelectComment(t *testing.T) {
	expectedComment := testdata.CommentTestData[0]
	got, err := repositories.SelectComment(testDB, expectedComment.CommentID)
	if err != nil {
		t.Fatal(err)
	}

	if got.CommentID != expectedComment.CommentID {
		t.Errorf("CommentID: get %d but want %d\n", got.CommentID, expectedComment.CommentID)
	}
	if got.ArticleID != expectedComment.ArticleID {
		t.Errorf("UserID: get %d but want %d\n", got.ArticleID, expectedComment.ArticleID)
	}
	if got.Message != expectedComment.Message {
		t.Errorf("UserName: get %s but want %s\n", got.Message, expectedComment.Message)
	}
}

func TestSelectCommentList(t *testing.T) {
	articleID := 1
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d but got ID %d\n", articleID, comment.ArticleID)
		}
	}
}

func TestInsertComment(t *testing.T) {
	insertCommentTestData := testdata.CommentInsertTestData

	expectedCommentID := 3
	newComment, err := repositories.InsertComment(testDB, insertCommentTestData)
	if err != nil {
		t.Error(err)
	}
	if newComment.CommentID != expectedCommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where message = ?
		`
		testDB.Exec(sqlStr, insertCommentTestData.Message)
	})
}

func TestUpdateComment(t *testing.T) {
	updateCommentTestData := testdata.CommentUpdateTestData
	updateCommentID := 2

	err := repositories.UpdateComment(testDB, updateCommentTestData, updateCommentID)
	if err != nil {
		t.Fatal(err)
	}

	got, err := repositories.SelectComment(testDB, updateCommentID)
	if err != nil {
		t.Fatal(err)
	}

	if got.Message != updateCommentTestData.Message {
		t.Errorf("update comment message is expected %s but got %s\n", updateCommentTestData.Message, got.Message)
	}
}

func TestDelete(t *testing.T) {
	deleteCommentID := 1

	err := repositories.DeleteArticle(testDB, deleteCommentID)

	if err != nil {
		if err != nil {
			t.Errorf("DeleteComment returned an error: %v", err)
		}
	}
}
