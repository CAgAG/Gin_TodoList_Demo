package http_func

import (
	"TodoList_demo/app/gateway/rpc"
	"TodoList_demo/grpc_proto/pb"
	"TodoList_demo/pkg/utils"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func CreateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err == nil {
		// 从请求上下文中拿到 JWT的数据
		user, err := utils.GetContextUserInfo(ctx.Request.Context())
		if err != nil {
			logging.Info("Token 验证失败")
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "Token 验证失败"))
			return
		}
		// ======================================
		if req.Title == "" {
			// 绑定成功, 但【有时候数据为空】 ???
			uid, _ := ctx.GetPostForm("uid")
			uid_uint, _ := strconv.Atoi(uid)
			req.Uid = uint64(uid_uint)

			title, _ := ctx.GetPostForm("title")
			req.Title = title

			status_req, _ := ctx.GetPostForm("status")
			status_req_uint, _ := strconv.Atoi(status_req)
			req.Uid = uint64(status_req_uint)

			content, _ := ctx.GetPostForm("content")
			req.Content = content

			start_time_req, _ := ctx.GetPostForm("start_time")
			start_time_req_uint, _ := strconv.Atoi(start_time_req)
			req.StartTime = int64(start_time_req_uint)

			end_time_req, _ := ctx.GetPostForm("end_time")
			end_time_req_uint, _ := strconv.Atoi(end_time_req)
			req.StartTime = int64(end_time_req_uint)
		}
		// =====================================

		req.Uid = uint64(user.Id)
		taskRes, err := rpc.CreateTask(ctx, &req)
		if err != nil {
			logging.Info("Token 验证失败")
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "RPC 执行失败"))
			return
		}
		ctx.JSON(http.StatusOK, utils.RespSuccess(ctx, taskRes))

	} else {
		ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "任务序列化绑定失败"))
		return
	}
}

func UpdateTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err == nil {
		// 从请求上下文中拿到 JWT的数据
		user, err := utils.GetContextUserInfo(ctx.Request.Context())
		if err != nil {
			logging.Info("Token 验证失败")
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "Token 验证失败"))
			return
		}
		req.Uid = uint64(user.Id)
		taskRes, err := rpc.UpdateTask(ctx, &req)
		if err != nil {
			logging.Info("Token 验证失败")
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "RPC 执行失败"))
			return
		}
		ctx.JSON(http.StatusOK, utils.RespSuccess(ctx, taskRes))

	} else {
		ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "任务序列化绑定失败"))
		return
	}
}

func DelTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err == nil {
		// 从请求上下文中拿到 JWT的数据
		user, err := utils.GetContextUserInfo(ctx.Request.Context())
		if err != nil {
			logging.Info("Token 验证失败")
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "Token 验证失败"))
			return
		}
		req.Uid = uint64(user.Id)
		taskRes, err := rpc.DelTask(ctx, &req)
		if err != nil {
			logging.Info("Token 验证失败")
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "RPC 执行失败"))
			return
		}
		ctx.JSON(http.StatusOK, utils.RespSuccess(ctx, taskRes))

	} else {
		ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "任务序列化绑定失败"))
		return
	}
}

func GetTasksListHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err == nil {
		// 从请求上下文中拿到 JWT的数据
		user, err := utils.GetContextUserInfo(ctx.Request.Context())
		if err != nil {
			logging.Info("Token 验证失败")
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "Token 验证失败"))
			return
		}
		req.Uid = uint64(user.Id)
		taskRes, err := rpc.GetTasksList(ctx, &req)
		if err != nil {
			logging.Info("Token 验证失败")
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "RPC 执行失败"))
			return
		}
		ctx.JSON(http.StatusOK, utils.RespSuccess(ctx, taskRes))

	} else {
		ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "任务序列化绑定失败"))
		return
	}
}

func GetTaskHandler(ctx *gin.Context) {
	var req pb.TaskRequest
	if err := ctx.ShouldBind(&req); err == nil {
		// 从请求上下文中拿到 JWT的数据
		user, err := utils.GetContextUserInfo(ctx.Request.Context())
		if err != nil {
			logging.Info("Token 验证失败")
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "Token 验证失败"))
			return
		}
		req.Uid = uint64(user.Id)
		taskRes, err := rpc.GetTask(ctx, &req)
		if err != nil {
			logging.Info("Token 验证失败")
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "RPC 执行失败"))
			return
		}
		ctx.JSON(http.StatusOK, utils.RespSuccess(ctx, taskRes))

	} else {
		ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "任务序列化绑定失败"))
		return
	}
}
