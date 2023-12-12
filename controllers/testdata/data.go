package testdata

import "github.com/Naokiiiiiii/BlogApiPractice/models"

var articleTestData = []models.Article{
	models.Article{
		ID:          1,
		Title:       "firstPost",
		Contents:    "This is my first blog",
		UserID:      1,
		UserName:    "naoki",
		NiceList:    niceTestData,
		CommentList: commentTestData,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserID:   1,
		UserName: "naoki",
	},
}

var commentTestData = []models.Comment{
	models.Comment{
		CommentID: 1,
		ArticleID: 1,
		UserID:    1,
		UserName:  "naoki",
		Message:   "1st comment yeah",
	},
	models.Comment{
		CommentID: 2,
		ArticleID: 1,
		UserID:    1,
		UserName:  "naoki",
		Message:   "welcome",
	},
}

var niceTestData = []models.Nice{
	models.Nice{
		NiceID:    1,
		UserID:    1,
		ArticleID: 1,
	},
}
