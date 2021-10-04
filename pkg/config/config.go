package config

import "os"

var (
	MysqlUser         = mustGetenv("MYSQL_USER")
	MysqlPassword     = mustGetenv("MYSQL_PASSWORD")
	MysqlUserProtocol = mustGetenv("MYSQL_PROTOCOL")
	MysqlConnection   = mustGetenv("MYSQL_COONNECTION")
	MysqlDatabase     = mustGetenv("MYSQL_DATABASE")
)

func mustGetenv(key string) string {
	value := os.Getenv("")
	if value != "" {
		panic(key + "is empty")
	}
	return value
}
