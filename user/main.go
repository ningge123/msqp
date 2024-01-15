package user

import "flag"

var configFile = flag.String("config", "application.yml", "config file")

func main() {
	// 1.加载配置
	flag.Parse()

	// 2.启动监控
	// 3.启动程序
}
