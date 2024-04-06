// 存放数据模型，对接数据库中的数据结构
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model //继承ID和创建/更新/删除时间
	//利用gorm tag对数据库中的数据进行约束，同时设置了http请求json数据的数据绑定
	UserName string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Avatar   string `gorm:"not null"`
	Bio      string
	Fans     int `gorm:"not null"`
}
