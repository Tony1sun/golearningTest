package main

import (
	"context"
	"fmt"
	"go-micro-test/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
)

func main() {
	consulReg := consul.NewRegistry(registry.Addrs(":8500"))

	service := micro.NewService(
		micro.Name("greeter.client"),
		micro.Registry(consulReg),
	)

	service.Init()

	greeter := proto.NewGreeterService("greeter", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &proto.Request{Name: "World"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
