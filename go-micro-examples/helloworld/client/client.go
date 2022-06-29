package main

import (
	"context"
	"fmt"
	helloworld "go-micro-examples/helloworld/proto"

	"github.com/micro/go-micro/web"
)

func main() {
	// create a new service
	service := web.NewService()

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := helloworld.NewHelloworldService("go.micro.srv.HelloWorld", service.Client())

	// Make request
	rsp, err := cl.Call(context.Background(), &helloworld.Request{
		Name: "World!",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Msg)
}
