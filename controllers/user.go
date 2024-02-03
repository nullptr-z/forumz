package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullptr-z/forumz/dao"
	"github.com/nullptr-z/forumz/entity"
	"github.com/nullptr-z/forumz/logic"
	"go.uber.org/zap"
)

const ShouldBindByUser = "ShouldBind by user"

// @Summary get User by name
// @Tags 用户
// @Param name query string true "用户姓名"
// @Success 200 {string} json{code, success ,msg}
// @Router /user/getByName [get]
func GetByNameHandler(g *gin.Context) {
	// 获取参数
	param := g.Query("name")
	user, err := dao.QueryUserByName(param)
	if err != nil {
		zap.L().Info("QueryUserByName", zap.Error(err))
		// g.JSON(http.StatusOK, gin.H{"msg": []string{"Not found user", err.Error()}})
		ResponseError(g, CodeUserNotExists)
		return
	}
	// 响应结果
	g.JSON(http.StatusOK, gin.H{"success": true, "data": user})
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
	// 1.获取参数
	var repwd = g.PostForm("repassword")
	if err := g.ShouldBind(&user); err != nil {
		zap.L().Info(ShouldBindByUser, zap.Error(err))
		ResponseErrorWithMsg(g, CodeInvalidParams, []string{"Wrong parameters"})
		return
	}
	// 1.1校验参数
	if len(user.Password) == 0 || repwd != user.Password {
		// zap.L().Info("Password with repassword not equal")
		g.JSON(http.StatusBadRequest, gin.H{"msg": []string{"Password with repassword not equal"}})
		return
	}
	// 2.处理业务逻辑
	if err := logic.Register(&user); err != nil {
		zap.L().Error("Register failed", zap.Error(err))
		ResponseErrorWithMsg(g, CodeUserExists, []string{"Register failed user existed"})
		return
	}
	// 3.响应
	g.JSON(http.StatusOK, gin.H{"msg": "Register user success", "success": true})
}

// @Summary Login User
// @Tags 用户
// @Param name formData string true "用户姓名"
// @Param password formData string false "密码"
// @Success 200 {string} json{code, success ,msg}
// @Router /user/login [post]
func LoginHandler(g *gin.Context) {
	var user entity.User
	// 1.获取参数
	if err := g.ShouldBind(&user); err != nil {
		zap.L().Info(ShouldBindByUser, zap.Error(err))
		ResponseError(g, CodeInvalidPassword)
		return
	}
	// 1.1校验参数

	// 2.处理业务逻辑
	token, err := logic.Login(&user)
	if err != nil {
		zap.L().Error("Login failed", zap.String("username", user.Username), zap.Error(err))
		// g.JSON(http.StatusBadRequest, gin.H{"msg": []string{"Login failed", err.Error()}})
		ResponseError(g, CodeUserNotExists, "Login failed", err.Error())
		return
	}
	// 3.响应
	ResponseSuccess(g, map[string]string{"token": token}, "Login user", user.Username)
}
