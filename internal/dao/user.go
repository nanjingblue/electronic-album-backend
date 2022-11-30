package dao

import (
	"electronic-gallery/global"
	"electronic-gallery/internal/model"
)

type userDAO struct{}

// User 数据库操作对象 使用单例模式
var User *userDAO

func init() {
	User = &userDAO{}
}

// GetUserByUsername 根据 username 查找用户
func (u userDAO) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := global.DBEngine.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetUserByID 通过 id 获取 user
func (u userDAO) GetUserByID(id uint) (model.User, error) {
	var user model.User
	err := global.DBEngine.First(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser 创建新用户
func (u userDAO) CreateUser(user *model.User) error {
	return global.DBEngine.Create(&user).Error
}

// UpdateUser 更新用户
func (u userDAO) UpdateUser(user *model.User) error {
	return global.DBEngine.Save(&user).Error
}

// DeleteUserByUserID 删除用户
func (u userDAO) DeleteUserByUserID(uid uint) error {
	var user model.User
	return global.DBEngine.Delete(&user, uid).Error
}
