package main

import (
	"flag"
	"log"
	"shortlink/router"
	"shortlink/server"

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
//go:generate $ETCD_HOME/etcdctl --endpoints $ETCD_ENPOINTS put /shortLink/ '{"startNum":10, "endNum":20}' > /dev/null
func main() {
	server.Init(*conf)

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
