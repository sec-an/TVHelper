package mysql

import (
	"TVHelper/global"
	"database/sql"
	"fmt"
	"time"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	mysqlAddr := fmt.Sprintf("%s:%d", global.MysqlSetting.Host, global.MysqlSetting.Port)
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=true&loc=Local",
		global.MysqlSetting.Username, global.MysqlSetting.Password, mysqlAddr, global.MysqlSetting.Database, "utf8")

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		global.Logger.Fatal("MySQL Connect Failed!", zap.Error(err))
	}
	conn.SetConnMaxLifetime(7 * time.Second) //设置空闲时间，这个是比mysql 主动断开的时候短
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)
	return conn
}
