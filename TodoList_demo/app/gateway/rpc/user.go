package rpc

import (
	"TodoList_demo/grpc_proto/pb"
	"TodoList_demo/pkg/status"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// 绑定 RPC

func UserLogin(ctx *gin.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	resp, err = UserService.UserLogin(ctx, req)

	if resp == nil {
		logging.Info("User服务返回失败 == resp = nil")
		logging.Info(err)
		return
	}

	if err != nil || resp.Code != status.Success {
		resp.Code = status.Error
		return
	}
	return
}

func UserRegister(ctx *gin.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	resp, err = UserService.UserRegister(ctx, req)

	if resp == nil {
		logging.Info("User服务返回失败 == resp = nil")
		logging.Info(err)
		return
	}

	if err != nil || resp.Code != status.Success {
		resp.Code = status.Error
		logging.Info("User服务返回失败")
		logging.Info(err)
		return
	}
	return
}
