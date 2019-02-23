package main

import (
	"context"
	"fmt"
	"github.com/kaiouz/GoShare/config"
	"github.com/kaiouz/GoShare/res"
	"github.com/kaiouz/GoShare/sd"
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

	// 启动服务发现
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go sd.StartSD(ctx, config.Config.Port)

	if err := server.Start(); err != nil {
		fmt.Println(err)
	}
}
