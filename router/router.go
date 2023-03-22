package router

import (
	"Ygo/pkg/parser"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func RegisterRoute(s *gin.Engine) {
	s.Any("/*path", func(context *gin.Context) {
		response := parser.ParseRequest(context.Request)
		defer response.Body.Close()
		context.Status(response.StatusCode)
		for name, values := range response.Header {
			for _, value := range values {
				context.Header(name, value)
			}
		}
		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println("read response body failed!", err)
		}
		context.Writer.Write(data)
	})
}
