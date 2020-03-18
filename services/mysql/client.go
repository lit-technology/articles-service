package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/philip-bui/articles-service/config"
	"github.com/rs/zerolog/log"
)

const (
	NoResults = "sql: no rows in result set"
	DuplicateKeyPrefix = ""
)

var (
	// DB is the exported PostgreSQL Client.
	DB *sql.DB
)

func init() {
	LoadDB()
}

func LoadDB() {
	if DB != nil {
		return
	}
	db, err := sql.Open("mysql", config.MySQLUser+":"+config.MySQLPass+"@/"+config.MySQLDB)
	if err != nil {
		log.Fatal().Err(err).Msg("error connecting to postgres dw")
	}
	db.SetMaxIdleConns(0)
	DB = db
}
