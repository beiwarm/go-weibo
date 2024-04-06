package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"weibo/base64"
	"weibo/common"
	"weibo/dto"
	"weibo/model"
	"weibo/response"
)

// 根据给定的用户名获取该用户的头像
func Avatar(ctx *gin.Context) {
	paramUserName := ctx.Params.ByName("UserName")
	//查询用户名是否存在
	exist, userData := common.UserExist(paramUserName)
	if !exist {
		response.Fail(ctx, http.StatusForbidden, "用户名不存在")
		return
	}
	response.SuccessWithData(ctx, "查询成功", gin.H{
		"Avatar": userData.Avatar,
	})
}

// 获取用户可展示的所有信息
func Show(ctx *gin.Context) {
	var temp model.User
	paramName := ctx.Params.ByName("UserName")
	result := common.DB.Where("UserName = ?", paramName).First(&temp)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Fail(ctx, http.StatusForbidden, "用户名不存在")
		return
	}
	response.SuccessWithData(ctx, "查询成功", gin.H{
		"avatar": temp.Avatar,
		"bio":    temp.Bio,
		"fans":   strconv.Itoa(temp.Fans),
	})
}

// 只有登录后才可以调用这个API
func VerifyPassword(ctx *gin.Context) {
	var param dto.User
	if !common.BindJSON(ctx, &param) {
		return
	}
	currentUser, _ := ctx.Get("CurrentUser")
	param.UserName = currentUser.(string)
	//获取当前用户的用户信息
	_, userData := common.UserExist(param.UserName)
	//验证密码是否正确
	if base64.Decode(userData.Password) != param.Password {
		response.Fail(ctx, http.StatusForbidden, "密码错误")
		return
	}
	response.Success(ctx, "密码正确")
}
