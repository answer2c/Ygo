package db

import (
	"Ygo/conf"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Conn *sql.DB

func InitConnectPool() {
	var err error
	Conn, err = sql.Open(conf.DBType, conf.DSN[conf.DBType])
	if err != nil {
		log.Fatal("create postgre connect error!", err)
		panic(err)
	}
	Conn.SetMaxIdleConns(conf.MaxIdleConn)
	Conn.SetConnMaxLifetime(conf.MaxConnLifeTime)
	Conn.Ping()
}

// QueryRows 通用的执行sql查询的方法
func QueryRows(sql string, param ...any) ([]map[string]interface{}, error) {
	stmt, err := Conn.Prepare(sql)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(param...)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			return
		}
	}()

	columns, err := rows.Columns()
	resultSet := make([]map[string]interface{}, 0)
	var values = make([]interface{}, len(columns))
	for k, _ := range values {
		values[k] = new(interface{})
	}
	for rows.Next() {
		if err = rows.Scan(values...); err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for k, v := range columns {
			var rawValue = *(values[k].(*interface{}))
			if b, ok := rawValue.([]byte); ok {
				row[v] = string(b)
			} else {
				row[v] = rawValue
			}
		}
		resultSet = append(resultSet, row)
	}
	return resultSet, nil
}
