package http_func

import (
	"TodoList_demo/app/gateway/rpc"
	"TodoList_demo/grpc_proto/pb"
	"TodoList_demo/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegisterHandler(ctx *gin.Context) {
	var req pb.UserRequest

	// 使用 postman 的 Body - raw, 不要用 form-data
	if err := ctx.ShouldBind(&req); err == nil {
		if req.UserName == "" {
			// 绑定成功, 但【有时候数据为空】 ???
			user_name, _ := ctx.GetPostForm("user_name")
			req.UserName = user_name
			password, _ := ctx.GetPostForm("password")
			req.Password = password
			password_confirm, _ := ctx.GetPostForm("password_confirm")
			req.PasswordConfirm = password_confirm
		}

		resp, err := rpc.UserRegister(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "注册RPC错误"))
			return
		}

		ctx.JSON(http.StatusOK, utils.RespSuccess(ctx, resp))
	} else {
		ctx.JSON(http.StatusBadRequest, utils.RespError(ctx, err, "数据绑定失败"))
	}
}

func UserLoginHandler(ctx *gin.Context) {
	var req pb.UserRequest

	if err := ctx.ShouldBind(&req); err == nil {
		if req.UserName == "" {
			// 绑定成功, 但【有时候数据为空】 ???
			user_name, _ := ctx.GetPostForm("user_name")
			req.UserName = user_name
			password, _ := ctx.GetPostForm("password")
			req.Password = password
		}

		resp, err := rpc.UserLogin(ctx, &req)
		if err != nil {
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "登录RPC错误"))
			return
		}

		// 返回 token
		token, err := utils.GenerateToken(uint(resp.UserDetail.Id))
		if err != nil {
			ctx.JSON(http.StatusOK, utils.RespError(ctx, err, "token 生成失败"))
		} else {
			ret := &utils.TokenData{Data: req, Token: token}
			ctx.JSON(http.StatusOK, utils.RespSuccess(ctx, ret))
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, utils.RespError(ctx, err, "数据绑定失败"))
	}
}
