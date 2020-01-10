package initRouter

import (
	"gin/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine  {
	router := gin.Default()
	// 加载静态资源
	router.Static("/statics","./statics")
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.Static("/avatar", "./avatar")
	// 全局加载模板
	router.LoadHTMLGlob("templates/*")
	indexGroup := router.Group("/")
	{
		indexGroup.GET("", handler.Index)
	}

	userGroup := router.Group("/user")
	{
		// 用户注册
		userGroup.POST("/register",handler.UserRegister)
		// 用户登入
		userGroup.POST("/login",handler.UserLogin)
		// 用户资料
		userGroup.GET("/profile",handler.UserProfile)
		// 用户更新
		userGroup.POST("/update",handler.UpdateUserProfile)
	}
	return router

}