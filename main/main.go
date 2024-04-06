package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"weibo/common"
	"weibo/controller/comment"
	"weibo/controller/file"
	"weibo/controller/like"
	"weibo/controller/tweet"
	"weibo/controller/user"
	"weibo/middleware"
)

func main() {
	//gorm使用了一个连接池来管理数据库连接，无需手动释放db
	//初始化数据库，若数据库出问题直接panic
	err := common.InitDB()
	if err != nil {
		panic(err)
	}
	//存储图片的静态文件夹路径
	images := file.DefaultPath
	err = common.InitFolder(images)
	if err != nil {
		panic(err)
	}
	//创建Gin框架核心引擎
	app := gin.Default()
	//配置静态文件服务
	app.Static("/static/images", images)
	app.Use(middleware.CORS())
	//添加路由
	api := app.Group("/api")
	userRoutes := api.Group("/user")
	userRoutes.POST("/register", user.Register)
	userRoutes.POST("/login", user.Login)
	userRoutes.GET("/auth", middleware.VerifyToken(), user.Auth)
	userRoutes.GET("/avatar/:UserName", user.Avatar)
	userRoutes.PUT("/avatar", middleware.VerifyToken(), user.ModifyAvatar)
	userRoutes.PUT("/password", middleware.VerifyToken(), user.ModifyPassword)
	userRoutes.GET("/:UserName", user.Show)
	userRoutes.POST("/password", middleware.VerifyToken(), user.VerifyPassword)
	fileRoutes := api.Group("/file")
	fileRoutes.POST("/upload", file.Upload)
	tweetRoutes := api.Group("/tweet")
	tweetRoutes.POST("", middleware.VerifyToken(), tweet.Create)
	//tweetRoutes.PUT(":ID", middleware.VerifyToken(), tweet.Update)
	//tweetRoutes.DELETE(":ID", middleware.VerifyToken(), tweet.Delete)
	tweetRoutes.GET("/byid/:ID", tweet.ShowByID)
	tweetRoutes.GET("/byname/:UserName", tweet.ShowByUserName)
	tweetRoutes.POST("/islike", like.IsLiking)
	tweetRoutes.POST("/like", like.Like)
	tweetRoutes.POST("/dislike", like.Dislike)
	tweetRoutes.POST("/comment", comment.Add)
	tweetRoutes.PUT("/comment", comment.Get)
	panic(app.Run(":5555"))
	//todo:CORS
	//todo:token相关
}
