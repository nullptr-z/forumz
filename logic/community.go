package logic

import (
	"github.com/nullptr-z/forumz/dao"
	"github.com/nullptr-z/forumz/entity"
)

func CommunityList() ([]entity.Community, error) {
	community, err := dao.CommunityList()
	return community, err
}

func CommunityFindById(id any) (entity.Community, error) {
	community, err := dao.CommunityFindById(id)
	return community, err
}
