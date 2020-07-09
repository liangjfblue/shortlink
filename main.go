package main

import (
	"flag"
	"log"
	"runtime/debug"
	"shortlink/config"
	"shortlink/models/db"
	"shortlink/models/redis"
	"shortlink/pkg"
	"shortlink/router"
	"shortlink/server"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	conf = flag.String("c", "./app.yml", "set configure file path")
	//conf = flag.String("c", "H:\\go_home\\opensource\\shortlink\\app.yml", "set configure file path")
)

func init() {
	flag.Parse()
}

//go:generate protoc -I ./rpc/proto --go_out=plugins=grpc:./rpc/proto ./rpc/proto/api.proto
func main() {
	//init config
	config.Init(*conf)

	//init log
	pkg.InitLog()

	//init db
	db.Init()

	//init redis
	hosts := strings.Split(config.Config().Redis.Host, ",")
	if err := redis.Init(hosts); err != nil {
		debug.PrintStack()
		panic(err)
	}

	//init router and start http server
	server.StartHttp(router.New())

	//init rpc service
	server.StartRpcServer()

	//ping and test http server health
	go func() {
		if err := server.PingServer(); err != nil {
			log.Fatal("http server start error: ", err)
		}
		logrus.Info("http server has been start ok")
	}()

	//init signal handle
	server.SignalHandle()
}
