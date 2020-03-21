package controllers

import (
	"testing"

	"github.com/philip-bui/articles-service/services/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetTagsByTagNameAndDate_DateFormat(t *testing.T) {
	date, err := mysql.ParseDateToMySQLFormat("20160922", GetTagsByTagNameAndDate_DateFormat)
	assert.NoError(t, err, "expected no error from parsing into sql date format")
	assert.Equal(t, "2016-09-22", date, "expected correct sql date format")
}
