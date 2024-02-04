package entity

type Vote struct {
	// UserID  直接从 gin的上下文能拿到
	PostId    string `json:"post_id" binding:"required"`               // 帖子 id
	Direction uint8  `json:"direction" binding:"required,oneof=0 1 2"` // 取消 0,赞成票 1，反对票 2,
}
