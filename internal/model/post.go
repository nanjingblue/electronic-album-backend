package model

import (
	"electronic-album/global"
	"github.com/jinzhu/gorm"
)

/*
Post 朋友圈发贴子
这里为了偷懒 规定贴子最多只能4张照片（twitter也是哦）
*/
type Post struct {
	gorm.Model
	UserID     uint `gorm:"not null"`
	User       User
	Content    string
	ImageOne   string
	ImageTwo   string
	ImageThree string
	ImageFour  string
}

func (p Post) GetAllPostByUserID(uid uint) ([]Post, error) {
	var posts []Post
	err := global.DBEngine.Where("user_id = ?", uid).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *Post) GetPostByID() error {
	return global.DBEngine.First(&p, p.UserID).Error
}

func (p *Post) CreatePost() error {
	return global.DBEngine.Create(&p).Error
}

func (p *Post) Update() error {
	return global.DBEngine.Update(&p).Error
}

func (p *Post) Delete() error {
	return global.DBEngine.Delete(&p).Error
}
