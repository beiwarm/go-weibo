package tweet

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"weibo/common"
	"weibo/model"
	"weibo/response"
)

func ShowByID(ctx *gin.Context) {
	var temp model.Tweet
	paramID := ctx.Params.ByName("ID")
	result := common.DB.Where("ID = ?", paramID).First(&temp)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		response.Fail(ctx, http.StatusForbidden, "推文不存在")
		return
	}
	response.SuccessWithData(ctx, "查询成功", gin.H{
		"userName":   temp.UserName,
		"content":    temp.Content,
		"image":      temp.Image,
		"tag":        temp.Tag,
		"createTime": temp.CreatedAt.Format("2006-01-02 15-04-05"),
	})
}

func ShowByUserName(ctx *gin.Context) {
	var temps []model.Tweet
	paramID := ctx.Params.ByName("UserName")
	common.DB.Where("UserName = ?", paramID).Find(&temps)
	var jsons []gin.H
	for i := 0; i < len(temps); i++ {
		jsons = append(jsons, gin.H{
			"id":         i + 1,
			"tweetID":    temps[i].ID,
			"userName":   temps[i].UserName,
			"content":    temps[i].Content,
			"image":      temps[i].Image,
			"tag":        temps[i].Tag,
			"createTime": temps[i].CreatedAt.Format("2006-01-02 15:04:05"),
			"like":       temps[i].Like,
		})
	}
	data, _ := json.Marshal(jsons)
	response.ResponseWithStr(ctx, http.StatusOK, "查询成功", string(data))
	//response.ResponseWithBytes(ctx, http.StatusOK, "查询成功", data)
}
