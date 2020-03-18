package config

import (
	"flag"

	"github.com/philip-bui/articles-service/pkg/env"
)

var (
	// Port of Service.
	Port string
	// MySQLHost is host of MySQL.
	MySQLHost string
	// MySQLPort is port of MySQL.
	MySQLPort string
	// MySQLUser is user to connect to MySQL.
	MySQLUser string
	// MySQLPass is password to connect to MySQL.
	MySQLPass string
	// MySQLDB is default MySQL Database to connect to.
	MySQLDB string
)

func init() {
	flag.StringVar(&Port, "port", "8080", "PORT")
	flag.StringVar(&MySQLHost, "mysqlhost", "127.0.0.1", "MYSQL_HOST")
	flag.StringVar(&MySQLPort, "mysqlport", "3306", "MYSQL_PORT")
	flag.StringVar(&MySQLUser, "mysqluser", "root", "MYSQL_USER")
	flag.StringVar(&MySQLPass, "mysqlpass", "password", "MYSQL_PASS")
	flag.StringVar(&MySQLDB, "mysqldb", "fairfax", "MYSQL_DB")
	flag.Parse()

	env.LoadEnv(&Port, "PORT")
	env.LoadEnv(&MySQLHost, "MYSQL_HOST")
	env.LoadEnv(&MySQLPort, "MYSQL_PORT")
	env.LoadEnv(&MySQLUser, "MYSQL_USER")
	env.LoadEnv(&MySQLPass, "MYSQL_PASS")
	env.LoadEnv(&MySQLDB, "MYSQL_DB")
}
