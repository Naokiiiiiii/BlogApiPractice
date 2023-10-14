package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

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
