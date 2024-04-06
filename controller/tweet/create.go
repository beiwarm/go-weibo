package tweet

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weibo/common"
	"weibo/controller/file"
	"weibo/dto"
	"weibo/model"
	"weibo/response"
)

func Create(ctx *gin.Context) {
	var param dto.Tweet
	if !common.BindJSON(ctx, &param) {
		return
	}
	//能到这个controller说明token验证过了
	temp, _ := ctx.Get("CurrentUser")
	currentUser := temp.(string)
	//不需要验证用户名存在了因为token机制保证了该用户名肯定存在
	////验证前端提供的用户名是否存在
	//if exist, _ := common.UserExist(param.UserName); !exist {
	//	response.Fail(ctx, http.StatusForbidden, "发布失败，用户名不存在")
	//	return
	//}

	//验证前端提供的推文是否为空
	if len(param.Content) == 0 {
		response.Fail(ctx, http.StatusForbidden, "内容为空")
		return
	}
	//用前端提供的参数创建推文对象
	var image string
	if param.Image == "" {
		image = ""
	} else {
		image = file.DefaultNetPath + param.Image
	}
	tweet := model.Tweet{
		UserName: currentUser,
		Content:  param.Content,
		Image:    image,
		Tag:      param.Tag,
	}
	//将创建好的对象序列化到数据库中
	result := common.DB.Create(&tweet)
	if result.Error != nil {
		response.Fail(ctx, http.StatusInternalServerError, "服务器错误，推文创建失败")
		return
	}
	if param.Tag != "" {
		//将tag添加进tag数据库中
		tagg := model.Tag{
			Content: param.Tag,
		}
		result = common.DB.Create(&tagg)
	}
	response.SuccessWithData(ctx, "发布成功", gin.H{
		"id": tweet.ID,
	})
}
