package model

import (
	"electronic-album/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

/*
Post 朋友圈发贴子
这里为了偷懒 规定贴子最多只能1张照片
*/
type Post struct {
	gorm.Model
	UserID  uint `gorm:"not null"`
	User    User
	Content string
	Image   string
}

func (p *Post) GetURl() string {
	client, _ := oss.New(global.OSSSetting.END_POINT, global.OSSSetting.ACCESS_KEY_ID, global.OSSSetting.ACCESS_KEY_SECRET)
	bucket, _ := client.Bucket(global.OSSSetting.BUCKET)
	signedGetURL, _ := bucket.SignURL(p.Image, oss.HTTPGet, 600)
	return signedGetURL
}

type PostSlice []Post

func (p PostSlice) Len() int {
	return len(p)
}

func (p PostSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PostSlice) Less(i, j int) bool {
	return p[i].CreatedAt.Before(p[j].CreatedAt)
}
