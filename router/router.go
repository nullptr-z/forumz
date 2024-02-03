package router

import (
	"github.com/nullptr-z/forumz/controllers/middleware"
	"github.com/nullptr-z/forumz/dao"
	_ "github.com/nullptr-z/forumz/docs"
	"github.com/nullptr-z/forumz/settings"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/nullptr-z/forumz/controllers"
	swge "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	if viper.Get("mode") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 发布模式
	}
	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery(), settings.LoggerFormateOutput)
	dao.InitializeDao()

	g.GET("/swagger/*any", swagger.WrapHandler(swge.Handler))
	g.GET("/auth", middleware.Authorization)
	// 用户
	userRouter := g.Group("/user")
	{
		userRouter.GET("/getByName", controllers.GetByNameHandler)
		userRouter.POST("/login", controllers.LoginHandler)
		userRouter.POST("/register", controllers.RegisterHandler)
	}
	// api/v1 版本
	v1 := g.Group("/api/v1")
	community := v1.Group("/community")
	{
		community.GET("/list", controllers.CommunityListHandler)
	}

	return g
}
