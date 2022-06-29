package main

import (
	"fmt"
	"go-gin-example/pkg/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
	}
}
