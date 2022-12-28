package dao

import (
	"electronic-gallery/global"
	"electronic-gallery/internal/model"
	"fmt"
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
func (f friendDao) FollowByUsername(user model.User, friendUsername string) error {
	user2, err := User.GetUserByUsername(friendUsername)
	if err != nil {
		return err
	}
	friend, err := f.GetFriend(user.ID, user2.ID)
	if err != nil {
		// 说明不存在记录，则创建
		friend.UserID = user.ID
		friend.FriendID = user2.ID
		friend.Relationship = model.Follow
		return global.DBEngine.Create(&friend).Error
	}
	if friend.Relationship == model.Follow {
		return fmt.Errorf("用户已关注")
	}
	friend.Relationship = model.Follow
	return nil
}

// BanByUsername 通过用户名拉黑
func (f friendDao) BanByUsername(user model.User, friendUsername string) error {
	user2, err := User.GetUserByUsername(friendUsername)
	if err != nil {
		return err
	}
	friend, err := f.GetFriend(user.ID, user2.ID)
	if err != nil {
		// 说明不存在记录，则创建
		friend.UserID = user.ID
		friend.FriendID = user2.ID
		friend.Relationship = model.UnFollow
		return global.DBEngine.Create(&friend).Error
	}
	if friend.Relationship == model.UnFollow {
		return fmt.Errorf("用户已拉黑")
	}
	friend.Relationship = model.UnFollow
	return nil
}

// GetFriend 根据 userID 和 friendID 获取 friend 记录，如果没有记录返回一个空记录
func (f friendDao) GetFriend(userID, friendID uint) (model.Friend, error) {
	friend := model.Friend{}
	err := global.DBEngine.Where("user_id = ? AND friend_id = ?", userID, friendID).First(&friend).Error
	if err != nil {
		return model.Friend{}, err
	}
	return friend, nil
}
