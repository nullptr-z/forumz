package dao

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
	"github.com/nullptr-z/forumz/dao/redi"
)

var (
	period       = float64(86400 * 7) // 一周
	scorePreVote = float64(432)       // 每票值多少分
)
var (
	ErrVoteTimeExpire = errors.New("expire vote time")
)

func PostVote(userId, postId string, dire float64) (err error) {
	// 从 Redis 取帖子发布的时间
	key := redi.GetVoteKey(redi.KeyPostTimeZSet)
	postTime := rds.ZScore(key, postId).Val()
	// 验证是否超过一周
	if float64(time.Now().Unix())-postTime > period {
		return ErrVoteTimeExpire
	}
	// 先查询当前用户之前为这个帖子投票的记录，重新计算分值
	preDire := rds.ZScore(redi.GetVoteKey(redi.KeyPostVoteSetZSetPref+postId), userId).Val()
	var op float64
	if dire > preDire {
		op = 1
	} else {
		op = 1
	}
	diff := math.Abs(preDire - dire)
	_, err = rds.ZIncrBy(redi.GetVoteKey(redi.KeyPostScoreZSet), scorePreVote*diff*op, postId).Result()
	if err != nil {
		return err
	}
	// 2.更新分数
	if dire == 0 {
		// 取消投票
		_, err = rds.ZRem(redi.GetVoteKey(redi.KeyPostVoteSetZSetPref+postId), postId).Result()
	}
	// 3.记录用户为该帖子投票的操作
	_, err = rds.ZAdd(redi.GetVoteKey(redi.KeyPostVoteSetZSetPref+postId), redis.Z{Score: dire, Member: userId}).Result()
	return
}
