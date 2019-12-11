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

var appDB = Config{}

func init() {
	if err := gotenv.Load(); err != nil {
		panic(err)
	}
	NewConfig()
}

func NewConfig() *Config {
	appDB = Config{
		MysqlHost:     os.Getenv("MYSQL_HOST"),
		MysqlPort:     os.Getenv("MYSQL_PORT"),
		MysqlUsername: os.Getenv("MYSQL_USERNAME"),
		MysqlPassword: os.Getenv("MYSQL_PASSWORD"),
		MysqlDB:       os.Getenv("MYSQL_DB"),
		AppHost:       os.Getenv("HOST"),
		AppPort:       os.Getenv("PORT"),
	}
	return &appDB
}

func GetConfig() Config{
	return appDB
}