package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weibo/model"
	"weibo/response"
)

// 验证用户是否存在，不存在返回false,存在返回true并返回该用户的数据模型
func UserExist(userName string) (bool, *model.User) {
	db := DB
	var temp model.User
	db.Where("UserName = ?", userName).First(&temp)
	//gormID是uint，且是从1开始自增的，所以0可以用来判断ID是否存在
	if temp.ID == 0 {
		//用户不存在
		return false, nil
	} else {
		//temp对象会通过栈逃逸分析创建在堆上，可以传出局部对象的引用给调用者
		return true, &temp
	}
}

// 验证并绑定JSON格式数据，格式有误时直接返回badRequest
func BindJSON(ctx *gin.Context, param any) bool {
	err := ctx.BindJSON(param)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "请使用正确的json格式发送请求")
		return false
	}
	return true
}
