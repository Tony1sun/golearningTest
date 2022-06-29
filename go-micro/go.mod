module example

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/client/http v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
)

replace github.com/micro/go-micro => github.com/Lofanmi/go-micro v1.16.1-0.20210804063523-68bbf601cfa4
