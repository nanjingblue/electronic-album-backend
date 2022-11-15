package model

import (
	"electronic-album/global"
	"github.com/jinzhu/gorm"
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
	UserID   uint `gorm:"not null"`
	User     User
	FriendID uint `gorm:"not null"`
}

// GetALLFriendsByUserID 通过 userID 获取好友列表
func (f Friend) GetALLFriendsByUserID(uid uint) ([]User, error) {
	var friends []User
	err := global.DBEngine.Where("user_id = ?", uid).First(&friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil
}

// FollowByUsername 通过 username 关注
func (f *Friend) FollowByUsername(username string) error {
	var friend User
	err := friend.GetUserByUsername(username)
	if err != nil {
		return err
	}
	f.FriendID = friend.ID
	return global.DBEngine.Create(&f).Error
}

// BanByUsername 通过用户名拉黑
func (f *Friend) BanByUsername(username string) error {
	var friend User
	err := friend.GetUserByUsername(username)
	if err != nil {
		return err
	}
	f.FriendID = friend.ID
	return global.DBEngine.Delete(&f).Error
}
