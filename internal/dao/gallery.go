package dao

import (
	"electronic-gallery/global"
	"electronic-gallery/internal/model"
)

type GalleryDAO struct{}

var Gallery *GalleryDAO

/*
init 相册的dao初始化
*/
func init() {
	Gallery = &GalleryDAO{}
}

// GetAllGalleryByUserID 根据 user_id 获取所有的 album
func (a GalleryDAO) GetAllGalleryByUserID(userID uint) ([]model.Gallery, error) {
	var albums []model.Gallery
	err := global.DBEngine.Where("user_id = ?", userID).Find(&albums).Error
	if err != nil {
		return nil, err
	}
	return albums, nil
}

/**
下面是对相册的增删改查操作
*/

// CreateGallery 新建相册
func (a GalleryDAO) CreateGallery(album model.Gallery) error {
	return global.DBEngine.Create(&album).Error
}

// DeleteGalleryByGalleryID 删除相册
func (a GalleryDAO) DeleteGalleryByGalleryID(albumID uint) error {
	var album model.Gallery
	return global.DBEngine.Delete(&album, albumID).Error
}

// UpdateGallery 更新相册
func (a GalleryDAO) UpdateGallery(album *model.Gallery) error {
	return global.DBEngine.Save(&album).Error
}

// GetGalleryByGalleryID 获取相册
func (a GalleryDAO) GetGalleryByGalleryID(galleryID uint) (model.Gallery, error) {
	var album model.Gallery
	err := global.DBEngine.First(&album, galleryID).Error
	if err != nil {
		return album, err
	}
	return album, nil
}

func (a GalleryDAO) GetRecycleByUserID(uid uint) (model.Gallery, error) {
	var recycle model.Gallery
	return recycle, global.DBEngine.Where("user_id = ? AND gallery_name = ?", uid, "回收站").First(&recycle).Error
}

func (a GalleryDAO) GetCollectionByUserID(uid uint) (model.Gallery, error) {
	var collection model.Gallery
	return collection, global.DBEngine.Where("user_id = ? AND gallery_name = ?", uid, "收藏夹").First(&collection).Error
}
