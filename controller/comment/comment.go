package comment

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"weibo/common"
	"weibo/dto"
	"weibo/model"
	"weibo/response"
)

func Add(ctx *gin.Context) {
	var param dto.Comment
	if !common.BindJSON(ctx, &param) {
		return
	}
	newCom := model.Comment{
		Content: param.Content,
		UserID:  param.UserID,
		ReplyID: param.ReplyID,
	}
	common.DB.Create(&newCom)
}

func Get(ctx *gin.Context) {
	var param dto.Comment
	if !common.BindJSON(ctx, &param) {
		return
	}
	//paramReplyID := ctx.Query("replyID")
	//log.Println("paramReplyID:" + paramReplyID)
	var data []model.Comment
	result := common.DB.Model(&model.Comment{}).Where("ReplyID = ?", param.ReplyID).Find(&data)
	//log.Printf("data:\n%v", result)
	var jsons []gin.H
	for i := 0; i < int(result.RowsAffected); i++ {
		jsons = append(jsons, gin.H{
			"id":      i + 1,
			"content": data[i].Content,
			"userID":  data[i].UserID,
		})
	}
	//log.Printf("json:\n%v", jsons)
	row, _ := json.Marshal(jsons)
	//log.Println("row")
	//log.Println(row)
	response.ResponseWithStr(ctx, http.StatusOK, "查询成功", string(row))
}
