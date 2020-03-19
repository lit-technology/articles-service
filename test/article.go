package test

import (
	"github.com/philip-bui/articles-service/models"
)

var Article = &models.Article{
	ID:    "1",
	Title: "latest science shows that potato chips are better for you than sugar",
	Date:  "2016-09-22",
	Body:  "some text, potentially containing simple markup about how potato chips are great",
	Tags:  []string{"health", "fitness", "science"},
}
