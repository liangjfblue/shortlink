/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"shortlink/config"
	"shortlink/rpc/proto"
	"shortlink/rpc/service"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StartHttp(g *gin.Engine) {
	go func() {
		if err := g.Run(fmt.Sprintf(":%d", config.Config().Server.Port)); err != nil {
			debug.PrintStack()
			panic(err)
		}
	}()
}

func StartRpcServer() {
	go func() {
		lis, err := net.Listen(config.Config().Server.Network, fmt.Sprintf(":%d", config.Config().Server.RpcPort))
		if err != nil {
			debug.PrintStack()
			panic(err)
		}

		s := grpc.NewServer()
		proto.RegisterShortLinkServer(s, &service.ShortLink{})

		if err := s.Serve(lis); err != nil {
			debug.PrintStack()
			panic(err)
		}
	}()
}

func SignalHandle() {
	logrus.Info("start http server")
	logrus.Info("start grpc server")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		logrus.Infof("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
			//return
		default:
			return
		}
	}
}

func PingServer() error {
	urlStr := fmt.Sprintf("http://127.0.0.1:%d%s", config.Config().Server.Port, config.Config().Server.PingUrl)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", urlStr, nil)
	if err != nil {
		panic(err)
	}

	for i := 0; i < config.Config().Server.PingTime; i++ {
		if resp, err := http.DefaultClient.Do(req); err == nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the http server")
}
