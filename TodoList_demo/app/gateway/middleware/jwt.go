package middleware

import (
	"TodoList_demo/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code uint32

		code = http.StatusOK
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "JWT Token为空",
			})
			ctx.Abort()
			return
		}
		claims, err := utils.CheckToken(token)
		if err != nil {
			code = http.StatusUnauthorized
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "JWT 无效Token",
			})
			ctx.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt.Unix() {
			code = http.StatusUnauthorized
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "JWT 权限过期, 请重新登陆",
			})
			ctx.Abort()
			return
		}
		// 将 JWT 拿到的json数据(id)放入 请求上下文
		// context 的 request 中添加 userKey
		ctx.Request = ctx.Request.WithContext(utils.NewContext(ctx.Request.Context(), &utils.UserInfo{Id: claims.Id}))
		ctx.Next()
	}
}
