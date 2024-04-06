package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"weibo/common"
	"weibo/model"
	"weibo/response"
	"weibo/uuid"
)

// 用于向服务器的静态文件目录上传文件
func Upload(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, "客户端未能正确上传文件")
		return
	}
	//为每一个文件生成一个独一无二的UUID
	uuid := uuid.New()
	//获取文件的扩展名
	extName := path.Ext(header.Filename)
	//服务器上存储的文件格式为Prefix+uuid.ext
	fileName := fmt.Sprintf("%s%s%s", Prefix, uuid, extName)
	//为每一个文件创建一个数据库项记录该文件的存在
	fileRecord := model.FileRecord{
		ID:         uuid,
		FileName:   fileName,
		ServerPath: DefaultPath + fileName,
		NetPath:    DefaultNetPath + fileName,
		FileExt:    strings.TrimPrefix(path.Ext(fileName), "."),
	}
	if err := common.DB.Create(&fileRecord).Error; err != nil {
		response.Fail(ctx, http.StatusInternalServerError, "无法在数据库中记录该文件")
		return
	}
	//创建空文件
	empty, err := os.Create(fileRecord.ServerPath)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, "服务器静态文件创建失败")
		return
	}
	//函数结束时自动关闭文件，若文件关闭失败报错
	defer func(empty *os.File) {
		err := empty.Close()
		if err != nil {
			response.Fail(ctx, http.StatusInternalServerError, "服务器静态文件关闭失败")
		}
	}(empty)
	//把上传的文件内容拷贝到创建的空文件中
	_, err = io.Copy(empty, file)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, "服务器静态文件复制失败")
		return
	}
	response.SuccessWithData(ctx, "文件上传成功", gin.H{
		"fileName": fileRecord.FileName,
		"filePath": fileRecord.NetPath,
	})
}
