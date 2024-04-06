package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 封装常用的http响应操作
func response(ctx *gin.Context, httpStatus int, msg string, data gin.H) {
	ctx.JSON(httpStatus, gin.H{
		"status": httpStatus,
		"msg":    msg,
		"data":   data,
	})
}

func ResponseWithStr(ctx *gin.Context, httpStatus int, msg string, data string) {
	ctx.JSON(httpStatus, gin.H{
		"status": httpStatus,
		"msg":    msg,
		"data":   data,
	})
}

func ResponseWithBytes(ctx *gin.Context, httpStatus int, msg string, data []byte) {
	ctx.JSON(httpStatus, gin.H{
		"status": httpStatus,
		"msg":    msg,
		"data":   data,
	})
}

func SuccessWithData(ctx *gin.Context, msg string, data gin.H) {
	response(ctx, http.StatusOK, msg, data)
}

func Success(ctx *gin.Context, msg string) {
	SuccessWithData(ctx, msg, nil)
}

func FailWithData(ctx *gin.Context, statusCode int, msg string, data gin.H) {
	response(ctx, statusCode, msg, data)
}

func Fail(ctx *gin.Context, statusCode int, msg string) {
	FailWithData(ctx, statusCode, msg, nil)
}
