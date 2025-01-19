package rpc

import (
	"TodoList_demo/grpc_proto/pb"
	"go-micro.dev/v4"
)

var (
	UserService pb.UserService
	TaskService pb.TaskService
)

func InitRPC() {
	// 网关是 客户端，下游的 user服务的是 服务端
	userMicroService := micro.NewService(micro.Name("userService.client"))
	userService := pb.NewUserService("rpcUserService", userMicroService.Client())
	UserService = userService

	// 网关是 客户端，下游的 task 服务的是 服务端
	taskMicroService := micro.NewService(micro.Name("taskService.client"))
	taskService := pb.NewTaskService("rpcTaskService", taskMicroService.Client())
	TaskService = taskService
}
