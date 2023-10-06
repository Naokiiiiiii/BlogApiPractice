package testdata

import "github.com/Naokiiiiiii/BlogApiPractice/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserID:   1,
		UserName: "naoki",
	},
	models.Article{
		ID:       2,
		Title:    "secondPost",
		Contents: "This is my second blog",
		UserID:   1,
		UserName: "naoki",
	},
}

var NiceTestData = []models.Nice{
	models.Nice{
		NiceID:    1,
		UserID:    1,
		ArticleID: 1,
	},
}
