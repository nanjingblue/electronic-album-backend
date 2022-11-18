package model

import (
	"electronic-album/global"
	"github.com/jinzhu/gorm"
)

type Gallery struct {
	gorm.Model
	GalleryName string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	User        User
	Cover       string
	Status      string `gorm:"default:'active';not null"`
	Description string
}

// GetAllAlbumByUserID 根据 user_id 获取所有的 album
func (a Gallery) GetAllAlbumByUserID(userID uint) ([]Gallery, error) {
	var albums []Gallery
	err := global.DBEngine.Where("user_id = ?", userID).Find(&albums).Error
	if err != nil {
		return nil, err
	}
	return albums, nil
}

/**
下面是对相册的增删改查操作
*/

// CreateAlbum 新建相册
func (a *Gallery) CreateAlbum() error {
	return global.DBEngine.Create(&a).Error
}

// DeleteAlbum 删除相册
func (a *Gallery) DeleteAlbum() error {
	return global.DBEngine.Delete(&a).Error
}

// UpdateAlbum 更新相册
func (a *Gallery) UpdateAlbum() error {
	return global.DBEngine.Update(&a).Error
}

// GetAlbum 获取相册
func (a *Gallery) GetAlbum() error {
	return global.DBEngine.First(&a, a.UserID).Error
}
