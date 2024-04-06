package tweet

//
//import (
//	"errors"
//	"github.com/gin-gonic/gin"
//	"gorm.io/gorm"
//	"net/http"
//	"weibo/common"
//	"weibo/dto"
//	"weibo/model"
//	"weibo/response"
//)
//
//func Update(ctx *gin.Context) {
//	var param dto.Tweet
//	if !common.BindJSON(ctx, &param) {
//		return
//	}
//	//验证前端提供的用户名是否存在
//	if exist, _ := common.UserExist(param.UserName); !exist {
//		response.Fail(ctx, http.StatusForbidden, "更新失败，用户名不存在")
//		return
//	}
//	//验证前端提供的推文是否为空
//	if len(param.Content) == 0 {
//		response.Fail(ctx, http.StatusForbidden, "更新失败，内容为空")
//		return
//	}
//	//验证前端提供的文章ID是否存在
//	var temp model.Tweet
//	paramID := ctx.Params.ByName("ID")
//	result := common.DB.Where("ID = ?", paramID).First(&temp)
//	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
//		response.Fail(ctx, http.StatusForbidden, "更新失败，推文不存在")
//		return
//	}
//	//验证前端提供的文章ID和用户名是否匹配
//	if temp.UserName != param.UserName {
//		response.Fail(ctx, http.StatusForbidden, "更新失败，推文与用户名不匹配")
//		return
//	}
//	//更新推文
//	result = common.DB.Model(&model.Tweet{}).Where("ID = ?", paramID).Update("Content", param.Content)
//	if result.Error != nil {
//		response.Fail(ctx, http.StatusInternalServerError, "更新失败，服务器错误")
//		return
//	}
//	response.Success(ctx, "更新成功")
//}
