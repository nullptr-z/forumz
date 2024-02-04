package redi

// redis key 尽量使用命名空间的方式，方便业务拆分、查询

const (
	Prefix                 = "forumz."
	KeyPostTimeZSet        = "post.time"   // zet;发帖时间
	KeyPostScoreZSet       = "post.score"  // zet;投票的分数
	KeyPostVoteSetZSetPref = "post.voted." // zet;投票的类型;参数是 post id
)

func GetVoteKey(key string) string {
	return Prefix + key
}
