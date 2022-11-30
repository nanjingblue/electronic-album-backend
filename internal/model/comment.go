package model

import "gorm.io/gorm"

/*
Comment
@PostID 对哪条贴子的评论
@UserID 留言者的ID
*/
type Comment struct {
	gorm.Model
	PostID  uint `gorm:"not null"`
	Post    Post
	Content string
	UserID  uint `gorm:"not null"`
	User    User
}
