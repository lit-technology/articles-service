package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/philip-bui/articles-service/models"
	"github.com/philip-bui/articles-service/services/mysql"
)

// TODO: Logging
func InsertArticle(w http.ResponseWriter, r *http.Request) {
	var a models.Article
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := mysql.InsertArticleAndArticleTags(&a); err != nil {
		if strings.HasPrefix(err.Error(), mysql.DuplicateKeyPrefix) {
			http.Error(w, "article "+a.ID+" has already been inserted", http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
}
