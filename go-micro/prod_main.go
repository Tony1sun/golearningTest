package main

import (
	"example/ProdService"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReq := consul.NewRegistry(
		registry.Addrs("localhost:8500"))

	ginRouter := gin.Default()
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST", "/prods", func(context *gin.Context) {
			context.JSON(200,
				gin.H{"data": ProdService.NewProdList(2)})
		})
	}

	server := web.NewService(
		web.Name("prodservice"),
		web.Address(":8001"),
		web.Handler(ginRouter),
		web.Registry(consulReq),
	)
	server.Init()
	server.Run()
}
