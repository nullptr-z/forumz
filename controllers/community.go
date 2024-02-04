package controllers

import (
	"errors"
	"fmt"

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

// @Summary Community Detail
// @Description query by id
// @Tags 社区
// @Router /api/v1/community/:id [get]
// @Success 200 {string} json{code, success ,msg}
func CommunityQueryByIdHandler(g *gin.Context) {
	id := g.Param("id")
	fmt.Println("ids:", id)
	if len(id) == 0 {
		errMsg := fmt.Sprint("id is:", id)
		zap.L().Error("logic.CommunityList() failed", zap.Error(errors.New(errMsg)))
		ResponseError(g, CodeInvalidParams, errMsg)
		return
	}
	// 查询所有社区，return [(id,name)]
	data, err := logic.CommunityFindById(id)
	if err != nil {
		zap.L().Error("logic.CommunityList() failed", zap.Error(err))
		ResponseError(g, CodeServerBusy)
		return
	}
	zap.L().Info("Response data:", zap.Any("result", data))
	ResponseSuccess(g, data, "Request CommunityList")
}
