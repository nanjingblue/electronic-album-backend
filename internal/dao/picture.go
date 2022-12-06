package dao

import (
	"electronic-gallery/global"
	"electronic-gallery/internal/model"
)

type PictureDao struct{}

var Picture *PictureDao

func init() {
	Picture = &PictureDao{}
}

// GetALLPicturesByGalleryID 根据相册 ID 获取所有的图片
func (a PictureDao) GetALLPicturesByGalleryID(galleryID uint) ([]model.Picture, error) {
	var pictures []model.Picture
	err := global.DBEngine.Where("gallery_id = ?", galleryID).Find(&pictures).Error
	if err != nil {
		return nil, err
	}
	return pictures, nil
}

func (p PictureDao) GetPictureByID(pID uint) (model.Picture, error) {
	var picture model.Picture
	return picture, global.DBEngine.First(&picture, pID).Error
}

// CreatePicture 新建图片
func (a PictureDao) CreatePicture(picture *model.Picture) error {
	return global.DBEngine.Create(&picture).Error
}

// DeletePicture 删除图片
func (a PictureDao) DeletePicture(picture *model.Picture) error {
	return global.DBEngine.Delete(&picture).Error
}

// DeletePictureByID 删除图片
func (a PictureDao) DeletePictureByID(id uint) error {
	return global.DBEngine.Delete(&model.Picture{}, id).Error
}

// UpdatePicture 更新图片
func (a PictureDao) UpdatePicture(picture *model.Picture) error {
	return global.DBEngine.Save(&picture).Error
}

// GetPicture 获取图片
//func (a PictureDao) GetPicture() error {
//	return global.DBEngine.First(&a, a.ID).Error
//}
