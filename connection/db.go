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

type Connector interface {
	GetConnection() *sqlx.DB
}

type Mysql struct {
}

func (ms Mysql) GetConnection() *sqlx.DB {
	config := config.GetConfig()
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
