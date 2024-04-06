package user

import (
	"github.com/gin-gonic/gin"
	"weibo/response"
)

// 测试token验证中间件
func Auth(ctx *gin.Context) {
	currUser, _ := ctx.Get("CurrentUser")
	response.SuccessWithData(ctx, "确认已授权", gin.H{
		"CurrentUser": currUser,
	})
}
