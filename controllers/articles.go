package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
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
	params := httprouter.ParamsFromContext(r.Context())
	ID := params.ByName("id")

	a, err := mysql.GetArticleByID(ID)
	if err != nil {
		if err.Error() == mysql.NoResults {
			http.Error(w, "article "+ID+" not found", 404)
		} else {
			http.Error(w, err.Error(), 500)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a)
}
