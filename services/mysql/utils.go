package mysql

import (
	"database/sql"
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	// TODO: Check if time library has MySQL Format.
	MySQLDateFormat = "2006-01-02"
)

func PrepareStatement(query string) *sql.Stmt {
	LoadDB()
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal().Err(err).Str("query", strings.Replace(query, "\t", "", 100)).Msg("error preparing statement")
	}
	return stmt
}
