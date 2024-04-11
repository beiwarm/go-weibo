package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"weibo/model"
)

var DB *gorm.DB

// 格式化参数并连接数据库，返回数据库对象的引用
func InitDB() error {
	const (
		user     = "root"
		password = "a@2admin"
		host     = "localhost"
		port     = "3306"
		database = "weibo"
	)
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		user, password, host, port, database)
	var err error
	DB, err = gorm.Open(mysql.Open(args), &gorm.Config{
		//防止gorm自动把列名转换为snake_case
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase: true,
		},
	})
	if err != nil {
		return fmt.Errorf("数据库连接失败: " + err.Error())
	}
	err = migrate()
	if err != nil {
		return fmt.Errorf("数据模型迁移失败: " + err.Error())
	}
	return nil
}

// 根据数据模型的类型在数据库中建表
// 若表已存在会补全数据模型中新增的字段，但不会删除已有的数据
func migrate() error {
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		return err
	}
	if err := DB.AutoMigrate(&model.Tweet{}); err != nil {
		return err
	}
	if err := DB.AutoMigrate(&model.TokenRecord{}); err != nil {
		return err
	}
	if err := DB.AutoMigrate(&model.FileRecord{}); err != nil {
		return err
	}
	if err := DB.AutoMigrate(&model.Tag{}); err != nil {
		return err
	}
	if err := DB.AutoMigrate(&model.Like{}); err != nil {
		return err
	}
	if err := DB.AutoMigrate(&model.Comment{}); err != nil {
		return err
	}
	if err := DB.AutoMigrate(&model.TagFollow{}); err != nil {
		return err
	}
	if err := DB.AutoMigrate(&model.UserFollow{}); err != nil {
		return err
	}
	return nil
}
