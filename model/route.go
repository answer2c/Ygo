package model

import (
	"Ygo/pkg/db"
	"Ygo/pkg/tool"
	"fmt"
)

const ROUTE_TABLENAME = "ygo_route"

const (
	ROUTE_STATUS_INVALID = iota
	ROUTE_STATUS_VALID
)

type Route struct {
	Id        int64
	Status    int8
	ServiceId int64
	Paths     []string
	Hosts     []string
	Methods   []string
}

func GetRoute(path string, host string, method string) (*Route, error) {
	sql := fmt.Sprintf("select id,service_id,paths from %s "+
		"where status = %d and $1 = any(paths) and $2 = any(hosts) and $3 = any(methods) limit 1",
		ROUTE_TABLENAME, ROUTE_STATUS_VALID)
	res, err := db.QueryRows(sql, path, host, method)
	if err != nil || len(res) == 0 {
		return nil, err
	}

	return convertRouteData(res[0]), nil
}

func convertRouteData(row map[string]interface{}) (res *Route) {
	res = new(Route)
	if id, ok := row["Id"]; ok {
		res.Id = id.(int64)
	}

	if serviceId, ok := row["service_id"]; ok {
		res.ServiceId = serviceId.(int64)
	}

	if status, ok := row["Status"]; ok {
		res.Status = int8(status.(int64))
	}

	if paths, ok := row["Paths"]; ok {
		res.Paths = tool.HandlePostgreArray(paths.(string))
	}

	if hosts, ok := row["Hosts"]; ok {
		res.Hosts = tool.HandlePostgreArray(hosts.(string))
	}

	if methods, ok := row["Methods"]; ok {
		res.Methods = tool.HandlePostgreArray(methods.(string))
	}
	return
}
