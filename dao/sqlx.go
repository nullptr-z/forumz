package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/nullptr-z/forumz/settings"
)

var sq *sqlx.DB

func InitializeDao() {
	sq = settings.GetDB()
}
