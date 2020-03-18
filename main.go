// fruits-service project fruits-service.go
package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/philip-bui/articles-service/config"
	"github.com/philip-bui/articles-service/controllers"
)

func main() {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/articles", controllers.InsertArticle)
	router.HandlerFunc(http.MethodGet, "/articles/:id", controllers.GetArticle)
	router.HandlerFunc(http.MethodGet, "/tags/:tagName/:date", controllers.GetTagsByTagNameAndDate)

	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
