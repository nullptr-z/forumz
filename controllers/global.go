package controllers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 全局上下文可访问字段, 后续上文件中可以 g.Get(key) 可以直接获取
var (
	ContextUserIDKey = "userId"
)

func GetUserId(g *gin.Context) string {
	id := g.MustGet(ContextUserIDKey)
	fmt.Println("id.(int64):", id.(int64))
	return strconv.FormatInt(id.(int64), 10)
}
