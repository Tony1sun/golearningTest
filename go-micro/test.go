package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
)

func callAPI2(s selector.Selector) {
	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myClient.NewRequest("prodservice", "/v1/prods", map[string]string{})
	var rsp map[string]interface{}
	err := myClient.Call(context.Background(), req, &rsp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rsp["data"])
}

// 原始调用方式
// func callAPI(addr string, path string, method string) (string, error) {
// 	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
// 	client := http.DefaultClient
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer res.Body.Close()
// 	buf, _ := ioutil.ReadAll(res.Body)
// 	return string(buf), nil
// }

func main() {
	consulReq := consul.NewRegistry(
		registry.Addrs("localhost:8500"),
	)
	mySelector := selector.NewSelector(
		selector.Registry(consulReq),
		selector.SetStrategy(selector.RoundRobin),
	)
	callAPI2(mySelector)
}
