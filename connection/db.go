package connection

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"loket-event-go/config"
	"sync"
)

var once sync.Once
var connection *sqlx.DB

func GetConnection(config *config.Config) *sqlx.DB {
	once.Do(func() {
		db, err := sqlx.Connect("mysql",
			fmt.Sprintf("%s:%s@(%s:%v)/%s?parseTime=true",
				config.MysqlUsername,
				config.MysqlPassword,
				config.MysqlHost,
				config.MysqlPort,
				config.MysqlDB))
		if err != nil {
			panic(err.Error())
		}
		connection = db
	})
	return connection
}
