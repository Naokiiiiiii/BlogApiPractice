package testdata

import "github.com/Naokiiiiiii/BlogApiPractice/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserID:   1,
		UserName: "naoki",
		NiceNum:  2,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserID:   2,
		UserName: "naoki",
		NiceNum:  4,
	},
}
