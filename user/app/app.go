package app

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"msqp/common/config"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) error {
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
				log.Println("user app quit")
				return nil
			case syscall.SIGHUP:
				stop()
				log.Println("user app quit")
				return nil
			default:
				return nil
			}
		}
	}
}
