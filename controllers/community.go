package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nullptr-z/forumz/logic"
	"go.uber.org/zap"
)

// @Summary Community List
// @Description no parameter
// @Tags 社区
// @Router /api/v1/community/list [get]
// @Success 200 {string} json{code, success ,msg}
func CommunityListHandler(g *gin.Context) {
	// 查询所有社区，return [(id,name)]
	data, err := logic.CommunityList()
	if err != nil {
		zap.L().Error("logic.CommunityList() failed", zap.Error(err))
		ResponseError(g, CodeServerBusy)
		return
	}
	ResponseSuccess(g, data, "Request CommunityList")
}
