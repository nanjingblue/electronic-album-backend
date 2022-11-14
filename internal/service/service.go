package service

import (
	"electronic-album/global"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Service struct {
	ctx     *gin.Context
	db      *gorm.DB
	session sessions.Session
}

func New(ctx *gin.Context) Service {
	svc := Service{ctx: ctx}
	svc.db = global.DBEngine
	svc.session = sessions.Default(ctx)
	return svc
}
