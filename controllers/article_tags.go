package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/philip-bui/articles-service/models"
	"github.com/philip-bui/articles-service/services/mysql"
)

const (
	GetTagsByTagNameAndDate_DateFormat = "20060102"
)

// TODO: Logging
func GetTagsByTagNameAndDate(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	tagName := params.ByName("tagName")
	date := params.ByName("date")
	if len(tagName) == 0 {
		http.Error(w, "invalid tagName", http.StatusBadRequest)
		return
	} else if len(date) == 0 {
		// TODO: Better date checking.
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	mySQLDate, err := mysql.ParseDateToMySQLFormat(date, GetTagsByTagNameAndDate_DateFormat)
	if err != nil {
		http.Error(w, "invalid date", http.StatusBadRequest)
		return
	}

	// TODO: Needs more work on differentiating error from not found.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	t, err := mysql.GetCountRecentArticlesAndRelatedTagsByTagNameAndDate(tagName, mySQLDate)
	if err != nil {
		// TODO: Differentiate between server error and no results.
		// http.Error(w, "internal server error", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&models.TagArticleMetadata{
			Tag:         tagName,
			Count:       0,
			Articles:    []string{},
			RelatedTags: []string{},
		})
		return
	}
	json.NewEncoder(w).Encode(t)
}
