package entity

type User struct {
	Id         uint   `form:"id"`
	UserId     uint   `form:"user_id"`
	Username   string `form:"name" json:"username" binding:"required"`
	Password   string `form:"password" json:"password" binding:"required"`
	Email      string
	Gender     Gender
	CreateTime uint
	UpdateTime uint
}

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
