package rpc

import (
	"TodoList_demo/grpc_proto/pb"
	"TodoList_demo/pkg/status"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateTask(ctx *gin.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.CreateTask(ctx, req)

	if resp == nil {
		logging.Info("Task服务返回失败 == resp = nil")
		logging.Info(err)
		return
	}

	if err != nil || resp.Code != status.Success {
		resp.Code = status.Error
		return
	}
	return
}

func UpdateTask(ctx *gin.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.UpdateTask(ctx, req)

	if resp == nil {
		logging.Info("Task服务返回失败 == resp = nil")
		logging.Info(err)
		return
	}

	if err != nil || resp.Code != status.Success {
		resp.Code = status.Error
		return
	}
	return
}

func DelTask(ctx *gin.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.DeleteTask(ctx, req)

	if resp == nil {
		logging.Info("Task服务返回失败 == resp = nil")
		logging.Info(err)
		return
	}

	if err != nil || resp.Code != status.Success {
		resp.Code = status.Error
		return
	}
	return
}

func GetTasksList(ctx *gin.Context, req *pb.TaskRequest) (resp *pb.TaskListResponse, err error) {
	resp, err = TaskService.GetTasksList(ctx, req)

	if resp == nil {
		logging.Info("Task服务返回失败 == resp = nil")
		logging.Info(err)
		return
	}

	if err != nil || resp.Code != status.Success {
		resp.Code = status.Error
		return
	}
	return
}

func GetTask(ctx *gin.Context, req *pb.TaskRequest) (resp *pb.TaskDetailResponse, err error) {
	resp, err = TaskService.GetTask(ctx, req)

	if resp == nil {
		logging.Info("Task服务返回失败 == resp = nil")
		logging.Info(err)
		return
	}

	if err != nil || resp.Code != status.Success {
		resp.Code = status.Error
		return
	}
	return
}
