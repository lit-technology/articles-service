package mysql

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/philip-bui/articles-service/models"
	"github.com/philip-bui/articles-service/test"
	"github.com/stretchr/testify/assert"
)

func TestGetCountRecentArticlesAndRelatedTagsByTagNameAndDate(t *testing.T) {
	date := "2016-07-22"
	for i := int64(5); i < 20; i++ {
		a := &models.Article{
			ID:    strconv.FormatInt(i, 10),
			Title: test.Article.Title,
			Date:  date,
			Body:  test.Article.Body,
			Tags:  test.Article.Tags,
		}
		if err := InsertArticleAndArticleTags(a); err != nil {
			if !strings.HasPrefix(err.Error(), DuplicateKeyPrefix) {
				t.Error("expected no error for inserting article " + err.Error())
			}
			time.Sleep(2 * time.Second)
		}
	}
	a := &models.Article{
		ID:    "20",
		Title: test.Article.Title,
		Date:  date,
		Body:  test.Article.Body,
		Tags:  []string{"art", "fitness", "sports"},
	}
	InsertArticleAndArticleTags(a)
	m, err := GetCountRecentArticlesAndRelatedTagsByTagNameAndDate("health", date)

	assert.NoError(t, err, "expected no error for get count recent articles and related tags by tag name and date")
	assert.Equal(t, "health", m.Tag)
	assert.Equal(t, 15, m.Count)
	assert.Equal(t, []string{"10", "11", "12", "13", "14", "15", "16", "17", "18", "19"}, m.Articles)
	assert.ElementsMatch(t, []string{"fitness", "science"}, m.RelatedTags)

	m, err = GetCountRecentArticlesAndRelatedTagsByTagNameAndDate("art", date)

	assert.NoError(t, err, "expected no error for get count recent articles and related tags by tag name and date")

	assert.Equal(t, "art", m.Tag)
	assert.Equal(t, 1, m.Count)
	assert.Equal(t, []string{"20"}, m.Articles)
	assert.ElementsMatch(t, []string{"fitness", "sports"}, m.RelatedTags)
}
