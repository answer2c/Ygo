package parser

import (
	"Ygo/model"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ParseRequest 解析请求并转发
func ParseRequest(req *http.Request) *http.Response {
	path := strings.Split(req.URL.Path, "/")
	if len(path) < 1 {
		log.Print("no path")
		return nil
	}
	r, err := model.GetRoute(path[1], req.Host, req.Method)
	if err != nil {
		return nil
	}

	if r == nil {

	}
	service, err := model.GetService(r.ServiceId)
	if err != nil {

	}

	url := service.Protocol + "://" + service.Host + ":" + strconv.Itoa(int(service.Port)) +
		"/" + service.Path + "?" + req.URL.RawQuery

	log.Println("request url:" + url)
	newReq, err := http.NewRequest(req.Method, url, nil)
	if err != nil {

	}
	resp, err := http.DefaultClient.Do(newReq)
	if err != nil {
		return nil
	}

	return resp
}
