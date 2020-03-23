package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/philip-bui/articles-service/models"
	"github.com/philip-bui/articles-service/test"
	"github.com/stretchr/testify/suite"
)

const (
	hostURL = "http://127.0.0.1:8080"
)

func TestMain(t *testing.M) {
	go func() {
		main()
	}()
	time.Sleep(1 * time.Second)
	t.Run()
}

func TestServer(t *testing.T) {
	suite.Run(t, new(MainSuite))
}

type MainSuite struct {
	suite.Suite
}

func (s *MainSuite) SetupSuite() {}

// func (s *MainSuite) TestHealthCheck() {
// 	resp, err := http.Get(hostURL + "/health")
// 	s.NoError(err)
// 	s.NotNil(resp)
// 	s.Equal(http.StatusOK, resp.StatusCode)
// }

func (s *MainSuite) TestPostArticle() {
	body, err := json.Marshal(test.Article)
	s.NoError(err)

	resp, err := http.Post(hostURL+"/articles", "application/json", bytes.NewBuffer(body))
	s.NoError(err)
	s.NotNil(resp)
	s.Contains([]int{http.StatusOK, http.StatusConflict}, resp.StatusCode)
}

func (s *MainSuite) TestGetArticle() {
	s.TestPostArticle()

	resp, err := http.Get(hostURL + "/articles/" + test.Article.ID)
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(http.StatusOK, resp.StatusCode)

	var a models.Article
	s.NoError(json.NewDecoder(resp.Body).Decode(&a))

	s.Equal(test.Article.ID, a.ID)
	s.Equal(test.Article.Title, a.Title)
	s.Equal(test.Article.Date, a.Date)
	s.Equal(test.Article.Body, a.Body)
	s.ElementsMatch(test.Article.Tags, a.Tags)
}

func (s *MainSuite) TestGetTagMetadataAtDate() {
	s.TestPostArticle()

	resp, err := http.Get(hostURL + "/tags/" + "health" + "/20160922")
	s.NoError(err)
	s.NotNil(resp)
	s.Equal(http.StatusOK, resp.StatusCode)

	var t models.TagArticleMetadata
	s.NoError(json.NewDecoder(resp.Body).Decode(&t))
	s.NotNil(t)

	s.Equal("health", t.Tag)
	s.True(t.Count > 0)
	s.True(len(t.Articles) > 0)
	s.True(len(t.RelatedTags) > 0)
	// TODO: Assert it works with missing data, correct values etc.
}
