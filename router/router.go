package router

import (
	_ "github.com/nullptr-z/forumz/docs"
	"github.com/nullptr-z/forumz/settings"

	"github.com/gin-gonic/gin"
	"github.com/nullptr-z/forumz/controllers"
	swge "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	g := gin.New()
	g.Use(gin.Logger(), gin.Recovery(), settings.LoggerFormateOutput)

	g.GET("/swagger/*any", swagger.WrapHandler(swge.Handler))
	userRouter := g.Group("/user")
	{
		userRouter.POST("/register", controllers.RegisterHandler)
		userRouter.GET("/getByName", controllers.GetByNameHandler)
		// userRouter.GET("/list", handler.GetUserList)
		// userRouter.GET("/create", handler.CreateUser)
		// userRouter.GET("/delete", handler.DeleteUser)
		// userRouter.POST("/update", handler.UpdateUser)
		// userRouter.GET("/findByName", handler.FindUserByName)
		// userRouter.POST("/login", handler.Login)
		// userRouter.GET("/sendMsg", handler.SendMessage)
		// userRouter.GET("/getFriendsListById", handler.SearchFriendListById)
		// userRouter.POST("/addFriend", handler.AddFriends)

		// userRouter.GET("/getById", handler.GetUserById)
		// userRouter.GET("/delete", handler.DeleteUser)
	}

	return g
}
