package main

import (
	"TodoList_demo/app/user/repository/db/dao"
	"TodoList_demo/app/user/service"
	"TodoList_demo/conf"
	"TodoList_demo/grpc_proto/pb"
	"fmt"
	logging "github.com/sirupsen/logrus"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	conf.Init()
	err := dao.NewMysqlDB()
	if err != nil {
		return
	}

	// etcd 注册
	// etcd是一个分布式的、高可用的、一致的key-value存储数据库
	// etcd 基于 gRPC 定义了清晰、面向用户的 API
	// etcd可集中管理配置信息，服务端将配置信息存储于etcd，客户端通过etcd得到服务配置信息，etcd监听配置信息的改变，发现改变通知客户端。
	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", conf.EtcdHost, conf.EtcdPort)))

	// 创建 微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(conf.UserServiceAddress),
		micro.Registry(etcdReg))
	microService.Init()
	// 绑定服务
	err = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserService())
	if err != nil {
		logging.Info("用户微服务绑定失败")
		logging.Info(err)
		return
	}
	err = microService.Run()
	if err != nil {
		logging.Info("用户微服务启动失败")
		logging.Info(err)
		return
	}
}
