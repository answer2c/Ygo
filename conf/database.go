package conf

import (
	"fmt"
	"time"
)

var DBPostgres = map[string]string{
	"host":   "127.0.0.1",
	"port":   "5431",
	"user":   "ybc",
	"passwd": "123",
}

var DSN = map[string]string{
	"postgres": fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DBPostgres["host"], DBPostgres["port"], DBPostgres["user"], DBPostgres["passwd"], "ygo"),
}

var DBType = "postgres" // 默认数据库类型  postgres

var MaxIdleConn = 5

var MaxConnLifeTime = 5 * time.Minute
