package logic

import (
	"github.com/nullptr-z/forumz/dao"
	"github.com/nullptr-z/forumz/entity"
)

/*
	帖子投票功能

- 简化版投票算法，投一票+432 分，86400/200 -> 200 张赞成票，为帖子延续一天
- direction含义: 1:2, 2:1, 1:0, 2:0

限制：发表后一周内可投票，一旦到期将 Redis中存入持久化DB，删除 KeyPostVoteSetZSetPref
*/
func PostVote(userId string, p *entity.Vote) error {
	return dao.PostVote(userId, p.PostId, float64(p.Direction))
}
