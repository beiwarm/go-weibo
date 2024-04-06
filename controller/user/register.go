package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weibo/base64"
	"weibo/common"
	"weibo/controller/file"
	"weibo/dto"
	"weibo/model"
	"weibo/response"
)

func Register(ctx *gin.Context) {
	////解析Http请求中的json信息并写入数据模型中，数据不合法时返回400错误
	var param dto.User
	if !common.BindJSON(ctx, &param) {
		return
	}
	//查询该用户是否存在
	if exist, _ := common.UserExist(param.UserName); exist {
		response.Fail(ctx, http.StatusForbidden, "用户名已存在")
		return
	}
	//若用户不存在，创建新用户并写入数据库
	hashedPassword := base64.Encode(param.Password)
	newUser := model.User{
		UserName: param.UserName,
		Password: hashedPassword,
		Avatar:   file.DefaultNetPath + file.DefaultAvatar,
		Bio:      param.Bio,
		Fans:     0,
	}
	common.DB.Create(&newUser)
	response.Success(ctx, "注册成功")
}
