package model

import (
	"gorm.io/gorm"
)

const (
	// UnFollow 未关注
	UnFollow string = "unfollow"
	// Follow 关注
	Follow string = "follow"
	// Ban 拉黑
	Ban string = "ban"
)

type Friend struct {
	gorm.Model
	UserID       uint `gorm:"not null"`
	User         User
	FriendID     uint   `gorm:"not null"`
	Relationship string `gorm:"default:'unfollow';not null"`
}
