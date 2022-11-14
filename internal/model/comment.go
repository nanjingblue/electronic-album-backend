package model

import "github.com/jinzhu/gorm"

type comment struct {
	gorm.Model
	UserID  uint
	User    User
	content string
}
