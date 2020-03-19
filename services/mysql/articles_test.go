package mysql

import (
	"strings"
	"testing"

	"github.com/philip-bui/articles-service/test"
	"github.com/stretchr/testify/assert"
)

func TestInsertArticleAndArticleTags(t *testing.T) {
	if err := InsertArticleAndArticleTags(test.Article); err != nil && !strings.HasPrefix(err.Error(), DuplicateKeyPrefix) {
		t.Error("expected no error for insert article " + err.Error())
	}
}

func TestGetArticleByID(t *testing.T) {
	a, err := GetArticleByID(test.Article.ID)
	assert.NoError(t, err, "expected no error for getting article by id")
	assert.ObjectsAreEqual(test.Article, a)
}
