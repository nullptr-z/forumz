package dao

import (
	"errors"

	"github.com/nullptr-z/forumz/entity"
)

func QueryUserByName(name string) (entity.User, error) {
	var user entity.User
	sqlStr := `select * from forum.user where username = $1`
	err := sq.Get(&user, sqlStr, name)
	return user, err
}

func CheckUserExist(name string) (err error) {
	sqlStr := `select count(1) from forum.user where username = $1`
	var count int
	err = sq.Get(&count, sqlStr, name)
	if count > 0 {
		return errors.New("用户已存在")
	}
	return err
}

func CreateUser(user *entity.User) error {
	sqlStr := `insert into forum.user(user_id,username,password) values($1,$2,$3)`
	_, err := sq.Exec(sqlStr, user.UserId, user.Username, user.Password)
	return err
}

// func LoginUser(name string) (entity.User, error) {
// 	var user entity.User
// 	sqlStr := `select * from forum.user where username = $1`
// 	err := sq.Get(&user, sqlStr, name)
// 	return user, err
// }
