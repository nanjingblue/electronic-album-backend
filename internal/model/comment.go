package model

import "github.com/jinzhu/gorm"

/*
comment
@PostID 对哪条贴子的评论
@UserID 留言者的ID
*/
type comment struct {
	gorm.Model
	PostID  uint `gorm:"not null"`
	Post    Post
	content string
	UserID  uint `gorm:"not null"`
	User    User
}
