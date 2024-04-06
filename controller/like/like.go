package like

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"weibo/common"
	"weibo/dto"
	"weibo/model"
	"weibo/response"
)

func IsLiking(ctx *gin.Context) {
	var param dto.Like
	if !common.BindJSON(ctx, &param) {
		return
	}
	var temp model.Like
	err := common.DB.Where("UserID = ? AND TweetID = ?", param.UserName, param.TweetID).First(&temp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.SuccessWithData(ctx, "", gin.H{
			"liking": false,
		})
	} else {
		response.SuccessWithData(ctx, "", gin.H{
			"liking": true,
		})
	}
}

func Like(ctx *gin.Context) {
	var param dto.Like
	if !common.BindJSON(ctx, &param) {
		return
	}
	newLike := model.Like{
		TweetID: param.TweetID,
		UserID:  param.UserName,
	}
	common.DB.Create(&newLike)
	common.DB.Model(&model.Tweet{}).Where("TweetID = ?", param.TweetID).Update("Like", param.OldLike+1)
	response.Success(ctx, "")
}

func Dislike(ctx *gin.Context) {
	var param dto.Like
	if !common.BindJSON(ctx, &param) {
		return
	}
	common.DB.Delete(&model.Like{}).Where("TweetID = ? AND UserID = ?", param.TweetID, param.UserName)
	common.DB.Model(&model.Tweet{}).Where("ID = ?", param.TweetID).Update("Like", param.OldLike-1)
	response.Success(ctx, "")
}
