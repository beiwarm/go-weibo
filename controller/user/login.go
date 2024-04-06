package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weibo/base64"
	"weibo/common"
	"weibo/dto"
	"weibo/response"
	"weibo/token"
)

func Login(ctx *gin.Context) {
	var param dto.User
	if !common.BindJSON(ctx, &param) {
		return
	}
	//验证用户名密码不为空
	if param.UserName == "" {
		response.Fail(ctx, http.StatusForbidden, "用户名为空")
		return
	}
	if param.Password == "" {
		response.Fail(ctx, http.StatusForbidden, "密码为空")
		return
	}
	//查询用户名是否存在
	exist, userData := common.UserExist(param.UserName)
	if !exist {
		response.Fail(ctx, http.StatusForbidden, "用户名不存在")
		return
	}
	//验证密码是否正确
	if base64.Decode(userData.Password) != param.Password {
		response.Fail(ctx, http.StatusForbidden, "密码错误")
		return
	}
	//生成token，由客户端进行保存，后续每次验证token进行用户鉴权
	tok, err := token.ReleaseToken(param)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, "服务器错误，生成token失败")
		return
	}
	response.SuccessWithData(ctx, "登录成功", gin.H{
		"token": tok,
	})
}
