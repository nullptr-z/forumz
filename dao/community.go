package dao

import (
	"github.com/nullptr-z/forumz/entity"
)

func CommunityList() ([]entity.Community, error) {
	var communities []entity.Community
	sqlStr := `SELECT * FROM forum.community;`
	// []使用 Select 而不是 Get
	err := sq.Select(&communities, sqlStr)
	if err != nil {
		return nil, err
	}
	return communities, nil
}

func CommunityFindById(id any) (entity.Community, error) {
	var communities entity.Community
	sqlStr := `SELECT * FROM forum.community where id = $1;`
	err := sq.Get(&communities, sqlStr, id)
	if err != nil {
		return communities, err
	}
	return communities, nil
}
