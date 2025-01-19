package main

import (
	"TodoList_demo/app/gateway/router"
	"TodoList_demo/app/gateway/rpc"
	"TodoList_demo/conf"
	"fmt"
	logging "github.com/sirupsen/logrus"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
	"time"
)

func main() {
	var err error

	conf.Init()
	rpc.InitRPC()

	// etcd 注册
	// etcd是一个分布式的、高可用的、一致的key-value存储数据库
	// etcd 基于 gRPC 定义了清晰、面向用户的 API
	// etcd可集中管理配置信息，服务端将配置信息存储于etcd，客户端通过etcd得到服务配置信息，etcd监听配置信息的改变，发现改变通知客户端。
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", conf.EtcdHost, conf.EtcdPort)))

	// 创建 微服务实例
	webService := web.NewService(
		web.Name("httpService"),
		web.Address(conf.GatewayServiceAddress),
		web.Registry(etcdReg),
		web.Handler(router.NewRouter()), // 注册
		web.RegisterTTL(time.Second*30),
		web.Metadata(map[string]string{"protocol": "http"}))
	err = webService.Init()
	if err != nil {
		logging.Info("网关微服务初始化失败")
		logging.Info(err)
		return
	}
	err = webService.Run()
	if err != nil {
		logging.Info("网关微服务启动失败")
		logging.Info(err)
		return
	}

}
