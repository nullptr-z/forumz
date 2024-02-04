package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nullptr-z/forumz/controllers"
	"github.com/nullptr-z/forumz/pkg/jwt"
)

func Authorization(g *gin.Context) {
	const bearerSchema = "Bearer "
	// 从请求头提取 TOken
	author := g.GetHeader("Authorization")
	if author == "" {
		controllers.ResponseError(g, controllers.CodeInvalidateAuth)
		g.Abort()
		return
	}
	// 分割取出 Token
	tokenString := strings.TrimPrefix(author, bearerSchema)
	if tokenString == author {
		controllers.ResponseError(g, controllers.CodeInvalidateAuth)
		g.Abort()
		return
	}

	// 解析验证 Token
	if claims, err := jwt.ParseTOken(tokenString); err != nil {
		controllers.ResponseError(g, controllers.CodeInvalidateToken)
	} else {
		// Token 是有效的
		fmt.Printf("解析token, User ID: %d, Username: %s\n", claims.UserId, claims.UserName)
		// 将用户信息添加到请求的上下文中
		g.Set(controllers.ContextUserIDKey, claims.UserId)
		g.Next() // 处理下一个请求
	}
}
