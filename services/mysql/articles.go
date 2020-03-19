package mysql

import (
	"strconv"
	"strings"

	"github.com/philip-bui/articles-service/models"
	"github.com/rs/zerolog/log"
)

const (
	InsertArticleSQL    = `INSERT INTO article(id, title, article_date, body) VALUES (?, ?, ?, ?)`
	InsertArticleTagSQL = `INSERT INTO article_tag(tag, article_date, article_id) VALUES (?, ?, ?)`
)

var (
	GetArticleByIDStmt = PrepareStatement(`
	SELECT A.id, A.title, A.article_date, A.body, GROUP_CONCAT(DISTINCT(T.tag)) AS tags
	FROM article AS A
	LEFT JOIN article_tag AS T
		ON T.article_id = A.id
	WHERE A.id = ?
	GROUP BY A.id
	`)
)

func InsertArticleAndArticleTags(a *models.Article) error {
	ID, err := strconv.Atoi(a.ID)
	if err != nil {
		log.Error().Err(err).Str("id", a.ID).Msg("error converting string to int")
		return err
	}

	tx, err := DB.Begin()
	if err != nil {
		log.Error().Err(err).Msg("error beginning transaction for insert article and tags")
		return err
	}

	if _, err := tx.Exec(InsertArticleSQL, a.ID, a.Title, a.Date, a.Body); err != nil {
		log.Error().Err(err).Int("id", ID).Str("title", a.Title).Str("article_date", a.Date).Str("body", a.Body).Msg("error inserting article")
		return err
	}

	insertArticleTagStmt, err := tx.Prepare(InsertArticleTagSQL)
	if err != nil {
		log.Error().Err(err).Str("query", InsertArticleTagSQL).Msg("error preparing statement for insert article tag")
		return err
	}

	for _, tag := range a.Tags {
		if _, err := insertArticleTagStmt.Exec(tag, a.Date, ID); err != nil {
			log.Error().Err(err).Str("tag", tag).Str("article_date", a.Date).Int("article_id", ID).Msg("error inserting article tag")
			return err
		}
	}

	if err := insertArticleTagStmt.Close(); err != nil {
		log.Error().Err(err).Msg("error closing statement for insert article tag")
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Error().Err(err).Msg("error committing transaction for insert article and tags")
		return err
	}
	return nil
}

func GetArticleByID(ID string) (*models.Article, error) {
	intID, err := strconv.Atoi(ID)
	if err != nil {
		log.Error().Err(err).Str("id", ID).Msg("error converting string to int")
		return nil, err
	}
	row := GetArticleByIDStmt.QueryRow(intID)

	a := &models.Article{}
	var tagsAgg string
	if err := row.Scan(&a.ID, &a.Title, &a.Date, &a.Body, &tagsAgg); err != nil {
		log.Error().Err(err).Msg("error scanning row for get article by id")
		return nil, err
	}
	a.Tags = strings.Split(tagsAgg, ",")
	return a, nil
}
