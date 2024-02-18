package app

import (
	"context"
	"google.golang.org/grpc"
	"msqp/common/config"
	"msqp/common/logs"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) error {
	logs.Init(config.Conf.AppName)

	server := grpc.NewServer()
	go func() {
		listen, err := net.Listen("tcp", config.Conf.Grpc.Addr)
		if err != nil {
			panic(err)
		}

		err = server.Serve(listen)
		if err != nil {
			panic(err)
		}
	}()

	stop := func() {
		server.Stop()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP)
	for {
		select {
		case <-ctx.Done():
			stop()
			return nil
		case s := <-c:
			switch s {
			case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
				stop()
				logs.Info("user app quit")
				return nil
			case syscall.SIGHUP:
				stop()
				logs.Info("user app quit")
				return nil
			default:
				return nil
			}
		}
	}
}
