package model

import (
	"electronic-album/global"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// Active 激活用户
	Active string = "active"
	// Inactive 未激活用户
	Inactive string = "inactive"
	// Suspend 被封禁用户
	Suspend string = "suspend"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null"`
	PasswordDigest string `gorm:"not null"`
	Status         string `gorm:"default:'active';not null"`
	Sex            string
	Age            uint
}

// GetUserByUsername 根据 username 查找用户
func (u *User) GetUserByUsername(username string) error {
	return global.DBEngine.Where("username = ?", username).First(&u).Error
}

// GetUserByID 通过 id 获取 user
func (u *User) GetUserByID(id uint) error {
	return global.DBEngine.First(&u, id).Error
}

// CreateUser 创建新用户
func (u *User) CreateUser() error {
	return global.DBEngine.Create(&u).Error
}

// UpdateUser 更新用户
func (u *User) UpdateUser() error {
	return global.DBEngine.Update(&u).Error
}

// DeleteUser 删除用户
func (u *User) DeleteUser() error {
	return global.DBEngine.Delete(&u).Error
}

// SetPassword 设置密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}
