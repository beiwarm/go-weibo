package token

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"weibo/base64"
	"weibo/common"
	"weibo/dto"
	"weibo/model"
)

func ReleaseToken(user dto.User) (string, error) {
	//查询该用户是否有token
	var record model.TokenRecord
	result := common.DB.Where("UserName = ?", user.UserName).First(&record)
	//如果没有找到属于该用户的token
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {

	} else if time.Now().Sub(record.CreatedAt) > expirationTime /*token过期了*/ {
		//删除该token记录
		common.DB.Delete(&record)
	} else {
		//找到有效token则直接返回，否则生成一个token
		return record.Token, nil
	}
	//base64加密用户名和当前时间戳作为用户的token
	token := base64.Encode(user.UserName + time.Now().Format("20060102150405"))
	//在数据库中记录用户的token
	result = common.DB.Create(&model.TokenRecord{
		UserName: user.UserName,
		Token:    token,
	})
	if result.Error != nil {
		return "", result.Error
	}
	return token, nil
}

const expirationTime = 72 * time.Hour

// Verify 验证token是否存在，若存在且有效则返回该token属于的用户名
func Verify(token string) (bool, string) {
	//ime.Now().Add(-expirationTime)
	var record model.TokenRecord
	result := common.DB.Where("Token = ?", token).First(&record)
	//如果没有找到属于该token
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, ""
	}
	//如果token时间已过期
	if time.Now().Sub(record.CreatedAt) > expirationTime {
		//删除该token记录
		common.DB.Delete(&record)
		return false, ""
	}
	return true, record.UserName
}
