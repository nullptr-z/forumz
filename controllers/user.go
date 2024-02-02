package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullptr-z/forumz/entity"
	"github.com/nullptr-z/forumz/logic"
	"go.uber.org/zap"
)

// @Summary get User by name
// @Tags 用户
// @Param name query string true "用户姓名"
// @Success 200 {string} json{code, success ,msg}
// @Router /user/getByName [get]
func GetByNameHandler(g *gin.Context) {
	// 获取参数
	param := g.PostForm("name")
	fmt.Println("param:", param)
	// 参数校验
	// 响应结果
	g.JSON(http.StatusOK, gin.H{"message": ""})
}

// @Summary register User
// @Tags 用户
// @Param name formData string true "用户姓名"
// @Param password formData string false "密码"
// @Param repassword formData string false "确认密码"
// @Success 200 {string} json{code, success ,msg}
// @Router /user/register [post]
func RegisterHandler(g *gin.Context) {
	var user entity.User
	// var salt = string(rand.Intn(10) * 1000)
	// 1.获取参数
	var repwd = g.PostForm("repassword")
	if err := g.ShouldBind(&user); err != nil || repwd != user.Password {
		zap.L().Error("ShouldBind to User", zap.Error(err))
		g.JSON(http.StatusBadRequest, gin.H{"msg": []string{"Wrong parameters", err.Error()}})
		return
	}
	// 1.1校验参数
	// if len(user.Password) == 0 || repwd != user.Password {
	// 	zap.L().Error("Validate failed password", zap.Strings("password", []string{user.Password, repwd}))
	// 	g.JSON(http.StatusBadRequest, gin.H{"msg": "Validate failed password"})
	// 	return
	// }
	// 2.处理业务逻辑
	logic.Register(&user)
	// 3.响应
	g.JSON(200, gin.H{"msg": "Register user success", "success": true})
}
