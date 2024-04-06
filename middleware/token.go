package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weibo/response"
	"weibo/token"
)

func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//局部变量会和包名重名
		t := ctx.Request.Header.Get("Authorization")
		exist, userName := token.Verify(t)
		//如果token不存在，则直接返回未授权的操作
		if !exist {
			response.Fail(ctx, http.StatusUnauthorized, "未授权的操作")
			//如果需要授权的操作授权验证未通过，直接abort该请求防止该请求继续被处理
			ctx.Abort()
			return
		}
		//如果token存在且有效，将该token属于的用户名写入CurrentUser
		//供后续所有controller使用
		ctx.Set("CurrentUser", userName)
		ctx.Next()
	}
}
