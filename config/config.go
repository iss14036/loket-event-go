package config

import (
	"github.com/subosito/gotenv"
	"os"
)

type (
	Config struct {
		MysqlHost     string
		MysqlPort     string
		MysqlUsername string
		MysqlPassword string
		MysqlDB       string
		AppHost       string
		AppPort       string
	}
)

func init() {
	gotenv.Load()
}

func NewConfig() *Config {
	return &Config{
		MysqlHost:     os.Getenv("MYSQL_HOST"),
		MysqlPort:     os.Getenv("MYSQL_PORT"),
		MysqlUsername: os.Getenv("MYSQL_USERNAME"),
		MysqlPassword: os.Getenv("MYSQL_PASSWORD"),
		MysqlDB:       os.Getenv("MYSQL_DB"),
		AppHost:       os.Getenv("HOST"),
		AppPort:       os.Getenv("PORT"),
	}
}
