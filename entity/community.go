package entity

import "time"

type Community struct {
	ID            int       `db:"id"`
	CommunityID   int       `db:"community_id"`
	CommunityName string    `db:"community_name"`
	Introduction  string    `db:"introduction"`
	CreateTime    time.Time `db:"create_time"`
	UpdateTime    time.Time `db:"update_time"`
}
