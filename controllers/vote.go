package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nullptr-z/forumz/entity"
	"github.com/nullptr-z/forumz/logic"
	"go.uber.org/zap"
)

func PostVoteHandler(g *gin.Context) {
	// 获取请求参数
	p := new(entity.Vote)
	if err := g.ShouldBindJSON(p); err != nil {
		zap.L().Error("g.ShouldBindJSON", zap.Error(err))
		ResponseError(g, CodeInvalidParams)
		return
	}
	userId := GetUserId(g)
	fmt.Println("userId:", userId)
	if len(userId) == 0 {
		ResponseError(g, CodeNeedLogin)
		return
	}
	if err := logic.PostVote(userId, p); err != nil {
		zap.L().Error("logic.PostVote", zap.Error(err), zap.String("userId", userId), zap.Any("vote data", *p))
		ResponseError(g, CodeInvalidParams)
		return
	}
	ResponseSuccess(g, nil, "Success voted")
}
