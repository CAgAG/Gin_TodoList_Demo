package router

import (
	"TodoList_demo/app/gateway/http_func"
	"TodoList_demo/app/gateway/middleware"
	"TodoList_demo/pkg/status"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()

	api_v1 := ginRouter.Group("/api/v1")
	{
		api_v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(status.Success, "success")
		})

		api_v1.POST("/user/register", http_func.UserRegisterHandler)
		api_v1.POST("/user/login", http_func.UserLoginHandler)

		// JWT 保留用户状态, 使用中间件
		authed := api_v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.GET("task/list", http_func.GetTasksListHandler)
			authed.GET("task/get", http_func.GetTaskHandler)

			authed.POST("task/create", http_func.CreateTaskHandler)
			authed.POST("task/update", http_func.UpdateTaskHandler)
			authed.POST("task/delete", http_func.DelTaskHandler)
		}
	}

	return ginRouter
}
