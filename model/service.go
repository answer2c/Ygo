package model

import (
	"Ygo/pkg/db"
	"fmt"
)

type Service struct {
	Id       int64
	Status   int8
	Name     string
	Host     string
	Port     int16
	Protocol string
	Path     string
}

const SERVICE_TABLENAME = "ygo_service"

const (
	SERVICE_STATUS_INVALID = iota
	SERVICE_STATUS_VALID
)

func GetService(id int64) (*Service, error) {

	sql := fmt.Sprintf("select * from %s "+
		"where status = %d and id = $1 limit 1",
		SERVICE_TABLENAME, SERVICE_STATUS_VALID)
	res, err := db.QueryRows(sql, id)
	if err != nil || len(res) == 0 {
		return nil, err
	}
	return convertServiceData(res[0]), nil
}

func convertServiceData(row map[string]interface{}) (res *Service) {
	res = new(Service)
	if id, ok := row["id"]; ok {
		res.Id = id.(int64)
	}

	if status, ok := row["status"]; ok {
		res.Status = int8(status.(int64))
	}

	if path, ok := row["path"]; ok {
		res.Path = path.(string)
	}

	if host, ok := row["host"]; ok {
		res.Host = host.(string)
	}

	if protocol, ok := row["protocol"]; ok {
		res.Protocol = protocol.(string)
	}

	if name, ok := row["name"]; ok {
		res.Name = name.(string)
	}

	if port, ok := row["port"]; ok {
		res.Port = int16(port.(int64))
	}

	return
}
