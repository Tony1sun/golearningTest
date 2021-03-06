package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
)

func main() {
	ginRouter := gin.Default()
	ginRouter.Handle("GET", "/", func(context *gin.Context) {
		data := make([]interface{}, 0)
		context.JSON(200, gin.H{
			"data": data,
		})
	})
	ginRouter.Handle("GET", "/news", func(context *gin.Context) {
		context.String(200, "news api")
	})
	server := web.NewService(
		web.Address(":8000"),
		web.Handler(ginRouter),
	)
	server.Run()
}
