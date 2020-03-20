package mysql

import (
	"strings"

	"github.com/philip-bui/articles-service/models"
	"github.com/rs/zerolog/log"
)

var (
	GetCountRecentArticlesAndRelatedTagsByTagNameAndDateStmt = PrepareStatement(`
	WITH
		T AS (
			SELECT article_id
			FROM article_tag
			WHERE tag = ?
			AND article_date = ?
		),
		A AS (
			SELECT A.id, A.created
			FROM article AS A
			INNER JOIN T
				ON A.id = T.article_id
		),
		A_RECENT AS (
			SELECT id
			FROM A
			ORDER BY created DESC
			LIMIT 10
		)
		SELECT COUNT(DISTINCT(A.id)) AS count,
		GROUP_CONCAT(DISTINCT(A_RECENT.id)) AS articles,
		(SELECT GROUP_CONCAT(DISTINCT(T_RELATED.tag))
				FROM article_tag AS T_RELATED
				INNER JOIN T
					ON T_RELATED.article_id = T.article_id
				WHERE T_RELATED.tag != ?) AS related_tags
		FROM A, A_RECENT;
	`)
)

func GetCountRecentArticlesAndRelatedTagsByTagNameAndDate(tagName string, date string) (*models.TagArticleMetadata, error) {
	tagArticlesMetadata := &models.TagArticleMetadata{
		Tag: tagName,
	}
	row := GetCountRecentArticlesAndRelatedTagsByTagNameAndDateStmt.QueryRow(tagName, date, tagName)
	var articles, relatedTags string
	if err := row.Scan(&tagArticlesMetadata.Count, &articles, &relatedTags); err != nil {
		log.Error().Err(err).Str("tagName", tagName).Str("date", date).Msg("error scanning row for get count recent articles and related tags by tag name and date")
		return nil, err
	}
	tagArticlesMetadata.Articles = strings.Split(articles, ",")
	tagArticlesMetadata.RelatedTags = strings.Split(relatedTags, ",")
	return tagArticlesMetadata, nil
}
