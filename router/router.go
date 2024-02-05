package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullptr-z/forumz/controllers"
	"github.com/nullptr-z/forumz/dao"
	_ "github.com/nullptr-z/forumz/docs"
	"github.com/nullptr-z/forumz/middleware"
	"github.com/nullptr-z/forumz/settings"
	"github.com/spf13/viper"
	swge "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	dao.InitializeDao()

	if viper.Get("mode") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 发布模式
	}
	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery(), settings.LoggerFormateOutput)

	// 跨域访问配置
	g.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			// 处理探测性请求
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	g.GET("/swagger/*any", swagger.WrapHandler(swge.Handler))
	g.GET("/auth", middleware.Authorization)
	// 用户
	userRouter := g.Group("/user")
	{
		userRouter.GET("/getByName", controllers.GetByNameHandler)
		userRouter.POST("/login", controllers.LoginHandler)
		userRouter.POST("/register", controllers.RegisterHandler)
	}
	// // api/v1 版本
	v1 := g.Group("/api/v1", middleware.Authorization)
	// 社区
	community := v1.Group("/community")
	{
		community.GET("/:id", controllers.CommunityQueryByIdHandler)
		community.GET("/list", controllers.CommunityListHandler)
	}
	// 投票
	post := v1.Group("/post")
	{
		post.POST("/vote", controllers.PostVoteHandler)
	}
	// 静态界面加载
	g.NoRoute(func(c *gin.Context) {
		// 检查请求路径是否为API路径，如果不是，尝试返回首页
		if !isAPIPath(c.Request.URL.Path) {
			c.File("./static" + c.Request.URL.Path) // 返回SPA的首页
		} else {
			// 可以返回404页面或者404 JSON响应
			c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		}
	})

	return g
}

func isAPIPath(url string) bool {
	fmt.Println("url:", url)

	return false
}
