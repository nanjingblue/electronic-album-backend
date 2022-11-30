package model

import (
	"gorm.io/gorm"
	"time"
)

type UserPost struct {
	UserID    uint `gorm:"primaryKey"`
	User      User
	PostID    uint `gorm:"primaryKey"`
	Post      Post
	Liked     bool `gorm:"default:false"`
	Collected bool `gorm:"default:false"`
	Commented bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
