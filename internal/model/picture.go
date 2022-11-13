package model

import (
	"electronic-album/global"
	"github.com/jinzhu/gorm"
)

/*
Picture 图片
这里只设置了 album_id 为外键 应为 album 也是唯一属于一个 user 的
*/
type Picture struct {
	gorm.Model
	PictureName string `gorm:"not null"`
	Link        string `gorm:"not null"`
	AlbumID     uint   `gorm:"not null"`
	Album       Album
}

// GetALLPicturesByAlbumID 根据相册 ID 获取所有的图片
func GetALLPicturesByAlbumID(albumID uint) ([]Picture, error) {
	var pictures []Picture
	err := global.DBEngine.Where("album_id = ?", albumID).Find(&pictures).Error
	if err != nil {
		return nil, err
	}
	return pictures, nil
}

// CreatePicture 新建图片
func (a *Picture) CreatePicture() error {
	return global.DBEngine.Create(&a).Error
}

// DeletePicture 删除图片
func (a *Picture) DeletePicture() error {
	return global.DBEngine.Delete(&a).Error
}

// UpdatePicture 更新图片
func (a *Picture) UpdatePicture() error {
	return global.DBEngine.Update(&a).Error
}

// GetPicture 获取图片
func (a *Picture) GetPicture() error {
	return global.DBEngine.First(&a, a.ID).Error
}
