package main

import (
	"fmt"
	"github.com/kaiouz/GoShare/config"
	"github.com/kaiouz/GoShare/res"
	"github.com/kaiouz/GoShare/service"
)

func main() {
	// 初始化配置
	err := config.InitConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 加载资源
	resources := res.CreateResources(config.Config.Dir)
	// 启动服务
	server := service.CreateServer(config.Config.Port, resources)
	if err := server.Start(); err != nil {
		fmt.Println(err)
	}
}
