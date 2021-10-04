package mysql

import (
	"database/sql"
	"fmt"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	datasourceName := fmt.Sprintf(
		"%s:%s@%s(%s)/%s?parseTime=true&loc=Asia%%2FTokyo&charset=utf8mb4",
		config.MysqlUser,
		config.MysqlPassword,
		config.MysqlUserProtocol,
		config.MysqlConnection,
		config.MysqlDatabase,
	)
	return sql.Open("mysql", datasourceName)
}
