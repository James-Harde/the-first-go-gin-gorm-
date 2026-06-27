package main

import (
	"exchangeapp/config"
	"fmt"
)

func main() {
	// 初始化配置文件
	config.InitConfig()
	fmt.Printf("配置文件内容: %v, %v\n", config.AppConfig.App.Name, config.AppConfig.App.Port)
	fmt.Printf("配置文件内容: %v, %v", config.AppConfig.Database.Host, config.AppConfig.Database.Port)
}
