package dao

import (
	"fmt"

	"github.com/nullptr-z/forumz/entity"
)

func CommunityList() ([]entity.Community, error) {
	var communities []entity.Community
	sqlStr := `SELECT * FROM forum.community;`
	// 使用 Select 而不是 Get
	err := sq.Select(&communities, sqlStr)
	if err != nil {
		fmt.Println("ereqwewqer:", err)
		return nil, err
	}
	return communities, nil
}
