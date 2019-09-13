package micro

//micro
var Config = `
{
  "consul" : "127.0.0.1:8500",
  "server_name":"dllgo.{{.Appname}}",
  "server_version":"latest"
}
`
var Maingo = `package main
import (
	"{{.Appname}}/cmd"
)
func main() {
	cmd.Run()
}
`
var Servergo = `package cmd
import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/service/grpc"
	"log"
	"time"
)
func Run() {
//加载配置项
	err := config.LoadFile("./conf/config.json")
	if err != nil {
		log.Fatalf("Could not load config file: %s", err.Error())
		return
	}
	conf := config.Map()
	//consul
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			conf["consul"].(string),
		}
	})
	// New Service
	service := grpc.NewService(
		micro.Name(conf["server_name"].(string)),
		micro.Version(conf["server_version"].(string)),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Registry(reg),
	)

	// Initialise service
	service.Init()
	
	// Register Handler

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
`
var Makefile = `
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/*.proto

.PHONY: build
build: proto

	go build -o {{.Appname}}-micro main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t {{.Appname}}-micro:latest
`
var Dockerfile = `
FROM alpine
ADD {{.Appname}}-micro /{{.Appname}}-micro
ENTRYPOINT [ "/{{.Appname}}-micro" ]
`
