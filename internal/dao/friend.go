package dao

import (
	"electronic-gallery/global"
	"electronic-gallery/internal/model"
)

type friendDao struct{}

var Friend *friendDao

func init() {
	Friend = &friendDao{}
}

// GetALLFollowingFriendsByUserID 通过 userID 获取好友列表
func (f friendDao) GetALLFollowingFriendsByUserID(uid uint) ([]model.User, error) {
	var friends []model.User
	err := global.DBEngine.Table("users").Select("*").Joins("join friends on users.id = friends.user_id").Scan(&friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil
}

// FollowByUsername 通过 username 关注
func (f friendDao) FollowByUsername(friend *model.Friend, username string) error {
	user, err := User.GetUserByUsername(username)
	if err != nil {
		return err
	}
	friend.FriendID = user.ID
	friend.Relationship = model.Follow
	return global.DBEngine.Create(&f).Error
}

// BanByUsername 通过用户名拉黑
func (f friendDao) BanByUsername(friend *model.Friend, username string) error {
	user, err := User.GetUserByUsername(username)
	if err != nil {
		return err
	}
	friend.FriendID = user.ID
	friend.Relationship = model.Ban
	return global.DBEngine.Delete(&friend).Error
}
