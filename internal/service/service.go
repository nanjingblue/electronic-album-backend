package service

import (
	"electronic-gallery/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	ctx *gin.Context
	db  *gorm.DB
}

func New(ctx *gin.Context) Service {
	svc := Service{ctx: ctx}
	svc.db = global.DBEngine
	return svc
}
