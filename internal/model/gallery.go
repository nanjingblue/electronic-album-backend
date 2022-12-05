package model

import (
	"electronic-gallery/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gorm.io/gorm"
	"strings"
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

func (g *Gallery) CoverURl() string {
	client, _ := oss.New(global.OSSSetting.END_POINT, global.OSSSetting.ACCESS_KEY_ID, global.OSSSetting.ACCESS_KEY_SECRET)
	bucket, _ := client.Bucket(global.OSSSetting.BUCKET)
	signedGetURL, _ := bucket.SignURL(g.Cover, oss.HTTPGet, 600)
	if global.OSSSetting.DOMAIN != "" {
		return strings.Replace(signedGetURL, global.OSSSetting.BUCKET + "." + global.OSSSetting.END_POINT, global.OSSSetting.DOMAIN, -1)
	}
	return signedGetURL
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
	return global.DBEngine.Save(&a).Error
}

// GetAlbum 获取相册
func (a *Gallery) GetAlbum() error {
	return global.DBEngine.First(&a, a.UserID).Error
}
