package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Naokiiiiiii/BlogApiPractice/models"
	"github.com/gorilla/mux"
)

func TestUpdateCommentHandler(t *testing.T) {
	var tests = []struct {
		name       string
		commentID  string
		resultCode int
	}{
		{name: "number pathparam", commentID: "1", resultCode: http.StatusOK},
		{name: "alphabet pathparam", commentID: "aaa", resultCode: http.StatusNotFound},
	}

	testComment := models.Comment{
		CommentID: 1,
		ArticleID: 1,
		UserID:    1,
		UserName:  "naoki",
		Message:   "1st comment yeah",
	}

	reqBodyBytes, _ := json.Marshal(testComment)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/comment/%s", tt.commentID)
			req := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(reqBodyBytes))

			res := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/comment/{id:[0-9]+}", cCon.UpdateCommentHandler).Methods(http.MethodPut)
			r.ServeHTTP(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}

func TestDeleteCommentHandler(t *testing.T) {
	var tests = []struct {
		name       string
		commentID  string
		resultCode int
	}{
		{name: "number pathparam", commentID: "1", resultCode: http.StatusOK},
		{name: "alphabet pathparam", commentID: "aaa", resultCode: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://localhost:8080/comment/%s", tt.commentID)
			req := httptest.NewRequest(http.MethodDelete, url, nil)

			res := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/comment/{id:[0-9]+}", cCon.DeleteCommentHandler).Methods(http.MethodDelete)
			r.ServeHTTP(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}
