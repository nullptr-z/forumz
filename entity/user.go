package entity

import (
	"database/sql"
	"time"
)

type User struct {
	Id         int64          `form:"id"`
	UserId     int64          `form:"userId" db:"user_id"`
	Username   string         `form:"name" json:"username" binding:"required"`
	Password   string         `form:"password" json:"password" binding:"required"`
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

// -- CREATE SCHEMA IF NOT EXISTS forum;
// CREATE TYPE forum.gender AS ENUM('unknown', 'man', 'woman');

// CREATE TABLE forum."user"(
//   id BIGSERIAL NOT NULL,
//   user_id BIGINT NOT NULL,
//   username VARCHAR(64) NOT NULL,
//   password VARCHAR(64) NOT NULL,
//   email VARCHAR(64) NOT NULL,
//   gender forum.gender NOT NULL DEFAULT 'unknown',
//   create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//   PRIMARY KEY (id),
//   UNIQUE (username),
//   UNIQUE (user_id)
