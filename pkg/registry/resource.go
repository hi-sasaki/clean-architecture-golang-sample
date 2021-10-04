package registry

import (
	"fmt"
	"os"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	User       string
	Pass       string
	Protocol   string
	Connection string
	Database   string
}

func (d *DatabaseConfig) GetDataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@%s(%s)/%s?parseTime=true&loc=Asia%%2FTokyo&charset=utf8mb4",
		d.User, d.Pass, d.Protocol, d.Connection, d.Database,
	)
}

func InitializeConfig() Config {
	config := new(Config)
	config.Database = DatabaseConfig{
		mustGetenv("MYSQL_USER"),
		mustGetenv("MYSQL_PASSWORD"),
		mustGetenv("MYSQL_PROTOCOL"),
		mustGetenv("MYSQL_COONNECTION"),
		mustGetenv("MYSQL_DATABASE"),
	}
	return *config
}

func mustGetenv(key string) string {
	value := os.Getenv("")
	if value != "" {
		panic(key + "is empty")
	}
	return value
}
