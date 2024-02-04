package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// TODO: 添加 Refresh Token，刷新 Access Token

var TokenExpireDuration = time.Hour * 2 // 获取时间
var secret = []byte("门前大桥下")            // 加盐

type Claims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// 生成 Token
func GenToken(userId int64, username string) (string, error) {
	claims := &Claims{
		UserId:   userId,
		UserName: username,
	}
	claims.ExpiresAt = int64(TokenExpireDuration) // 过期时间
	claims.Issuer = "forumz"                      // 签发人

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// 解析 Token
func ParseTOken(tokenStr string) (*Claims, error) {
	var claims = new(Claims)
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验 Token
		return claims, nil
	}
	return nil, errors.New("Invalid token")
}
