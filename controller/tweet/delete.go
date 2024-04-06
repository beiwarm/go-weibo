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
//func Delete(ctx *gin.Context) {
//	var param dto.Tweet
//	if !common.BindJSON(ctx, &param) {
//		return
//	}
//	//验证前端提供的文章ID是否存在
//	var temp model.Tweet
//	paramID := ctx.Params.ByName("ID")
//	result := common.DB.Where("ID = ?", paramID).First(&temp)
//	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
//		response.Fail(ctx, http.StatusForbidden, "删除失败，推文不存在")
//		return
//	}
//	//验证前端提供的用户名是否存在
//	if exist, _ := common.UserExist(param.UserName); !exist {
//		response.Fail(ctx, http.StatusForbidden, "删除失败，用户名不存在")
//		return
//	}
//	//验证前端提供的文章ID和用户名是否匹配
//	if temp.UserName != param.UserName {
//		response.Fail(ctx, http.StatusForbidden, "删除失败，推文不属于该用户")
//		return
//	}
//	// 删除文章
//	if err := common.DB.Delete(&temp).Error; err != nil {
//		response.Fail(ctx, http.StatusInternalServerError, "删除失败，服务器错误")
//		return
//	}
//	response.Success(ctx, "删除成功")
//}
