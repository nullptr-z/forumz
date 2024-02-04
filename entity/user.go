package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Id         int64          `form:"id"`
	UserId     int64          `form:"userId" db:"user_id"`
	Username   string         `form:"name" json:"username" binding:"required"`
	Password   string         `form:"password" json:"-" binding:"required"`
	Email      sql.NullString `form:"email" db:"email"`
	Gender     Gender         `form:"gender" db:"gender"`
	CreateTime time.Time      `db:"create_time"`
	UpdateTime time.Time      `db:"update_time"`
}

// ------form: http请求字段标签，db:数据库字段表情, binding: gin.Should_*_Bind标签

type Gender string

const (
	UNKNOWN Gender = "unknown"
	MAN     Gender = "man"
	WOMAN   Gender = "woman"
)
