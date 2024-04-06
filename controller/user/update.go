package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"weibo/base64"
	"weibo/common"
	"weibo/controller/file"
	"weibo/dto"
	"weibo/model"
	"weibo/response"
)

// 只有登录后才可以调用这个API
func ModifyAvatar(ctx *gin.Context) {
	currentUser, _ := ctx.Get("CurrentUser")
	var param dto.User
	if !common.BindJSON(ctx, &param) {
		return
	}

	//查询用户提供的文件名是否存在
	var tempFile model.FileRecord
	err := common.DB.Where("FileName = ?", param.Avatar).First(&tempFile).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Fail(ctx, http.StatusForbidden, "该文件不存在")
		return
	}
	//若文件名存在，将当前用户的头像修改为指定文件
	err = common.DB.Model(&model.User{}).Where("UserName = ?", currentUser).Update("Avatar", file.DefaultNetPath+param.Avatar).Error
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, "数据库出错")
		return
	}
	response.SuccessWithData(ctx, "修改成功", gin.H{
		"filePath": tempFile.NetPath,
	})
}

// 只有登录后才可以调用这个API
func ModifyPassword(ctx *gin.Context) {
	currentUser, _ := ctx.Get("CurrentUser")
	var param dto.User
	if !common.BindJSON(ctx, &param) {
		return
	}
	err := common.DB.Model(&model.User{}).Where("UserName = ?", currentUser).Update("Password",
		base64.Encode(param.Password)).Error
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, "服务器错误")
		return
	}
	response.Success(ctx, "修改成功")
}
