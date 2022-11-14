package dao

import (
	"electronic-album/global"
	"electronic-album/internal/model"
)

type AlbumDAO struct{}

var Album *AlbumDAO

func init() {
	Album = &AlbumDAO{}
}

// GetAllAlbumByUserID 根据 user_id 获取所有的 album
func (a AlbumDAO) GetAllAlbumByUserID(userID uint) ([]model.Album, error) {
	var albums []model.Album
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
func (a AlbumDAO) CreateAlbum(album *model.Album) error {
	return global.DBEngine.Create(&album).Error
}

// DeleteAlbumByAlbumID 删除相册
func (a AlbumDAO) DeleteAlbumByAlbumID(albumID uint) error {
	var album model.Album
	return global.DBEngine.Delete(&album, albumID).Error
}

// UpdateAlbum 更新相册
func (a *AlbumDAO) UpdateAlbum(album *model.Album) error {
	return global.DBEngine.Update(&album).Error
}

// GetAlbumByAlbumID 获取相册
func (a *AlbumDAO) GetAlbumByAlbumID(albumID uint) (*model.Album, error) {
	var album model.Album
	err := global.DBEngine.First(&album, albumID).Error
	if err != nil {
		return nil, err
	}
	return &album, nil
}
