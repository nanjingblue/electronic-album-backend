package model

import (
	"electronic-gallery/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gorm.io/gorm"
	"strings"
)

/*
Picture 图片
这里只设置了 album_id 为外键 应为 album 也是唯一属于一个 user 的
*/
type Picture struct {
	gorm.Model
	PictureName string `gorm:"not null"`
	Path        string `gorm:"not null"`
	UserID      uint   `gorm:"not null"` // 归属用户
	User        User
	GalleryID   uint `gorm:"not null"` // 归属相册
	Gallery     Gallery
}

func (p *Picture) CoverURl() string {
	client, _ := oss.New(global.OSSSetting.END_POINT, global.OSSSetting.ACCESS_KEY_ID, global.OSSSetting.ACCESS_KEY_SECRET)
	bucket, _ := client.Bucket(global.OSSSetting.BUCKET)
	signedGetURL, _ := bucket.SignURL(p.Path, oss.HTTPGet, 600)
	if global.OSSSetting.DOMAIN != "" {
		return strings.Replace(signedGetURL, global.OSSSetting.BUCKET + "." + global.OSSSetting.END_POINT, global.OSSSetting.DOMAIN, -1)
	}
	return signedGetURL
}
