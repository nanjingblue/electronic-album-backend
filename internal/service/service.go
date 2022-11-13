package service

import (
	"context"
	"electronic-album/global"
	"github.com/jinzhu/gorm"
)

type Service struct {
	ctx context.Context
	db  *gorm.DB
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.db = global.DBEngine
	return svc
}
