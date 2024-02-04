package dao

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/nullptr-z/forumz/settings"
)

var sq *sqlx.DB
var rds *redis.Client

func InitializeDao() {
	sq = settings.GetDB()
	rds = settings.GetRedis()
}
