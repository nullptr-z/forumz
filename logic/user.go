package logic

import (
	"errors"
	"fmt"

	"github.com/nullptr-z/forumz/dao"
	"github.com/nullptr-z/forumz/entity"
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

func Login(user *entity.User) (err error) {
	// 1.判断用户是否存在
	userf, err := dao.QueryUserByName(user.Username)
	if err != nil {
		// return errors.New(fmt.Sprint("Notfound user:", user.Username, err))
		return
	}

	// 验证密码
	isValid := utils.ValidPasswordDefaultSalt(user.Password, userf.Password)
	if !isValid {
		return errors.New(fmt.Sprint("Validate failed of password"))
	}
	// 构造实例
	// 写入数据库
	// dao.Login(user)
	return
}
