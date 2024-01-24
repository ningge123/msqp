package main

import (
	"context"
	"flag"
	"fmt"
	"msqp/common/config"
	"msqp/common/metrics"
	"msqp/user/app"
)

var configFile = flag.String("config", "application.yml", "config file")

func main() {
	// 1.加载配置
	flag.Parse()
	config.InitConfig(*configFile)
	// 2.启动监控
	go func() {
		err := metrics.Serve(fmt.Sprintf("0.0.0.0:%d", config.Conf.MetricPort))
		if err != nil {
			panic(err)
		}
	}()

	// 3.启动程序
	err := app.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
