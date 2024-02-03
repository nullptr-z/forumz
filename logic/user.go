package logic

import (
	"errors"
	"fmt"

	"github.com/nullptr-z/forumz/dao"
	"github.com/nullptr-z/forumz/entity"
	"github.com/nullptr-z/forumz/pkg/jwt"
	"github.com/nullptr-z/forumz/pkg/snowflake"
	"github.com/nullptr-z/forumz/utils"
)

func Register(user *entity.User) (err error) {
	// 1.判断用户是否存在
	err = dao.CheckUserExist(user.Username)
	if err != nil {
		return
	}

	// 2.生成 ID
	user.UserId = snowflake.GetId()
	// 密码加密
	user.Password = utils.CryptoPasswordDefaultSalt(user.Password)
	// 构造实例
	// 写入数据库
	dao.CreateUser(user)
	return
}

func Login(user *entity.User) (token string, err error) {
	// 1.查询用户
	userf, err := dao.LoginUser(user.Username)
	if err != nil {
		return
	}

	// 验证密码
	isValid := utils.ValidPasswordDefaultSalt(user.Password, userf.Password)
	if !isValid {
		return "", errors.New(fmt.Sprint("Validate failed of password"))
	}
	// 生成 TOken
	token, err = jwt.GenToken(userf.UserId, userf.Username)
	if err != nil {
		return
	}

	return
}
