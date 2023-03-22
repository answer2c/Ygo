package main

import (
	"Ygo/pkg/db"
	"Ygo/router"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	go func() {
		db.InitConnectPool()
	}()
}

func main() {
	server := gin.Default()
	router.RegisterRoute(server)
	err := server.Run(":8088")
	if err != nil {
		log.Fatal("run server failed!", err)
		return
	}
}
