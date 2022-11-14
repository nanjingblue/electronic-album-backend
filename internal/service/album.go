package service

import (
	"electronic-album/internal/model"
	"electronic-album/internal/serializer"
	"fmt"
)

type AlbumCreateRequest struct {
	AlbumName string `form:"album_name" json:"album_name" binding:"required,min=1,max=12"`
	UserID    uint   `form:"password" json:"user_id" binding:"required"`
}

func (acr *AlbumCreateRequest) Valid(userID uint) *serializer.Response {
	if acr.UserID != userID {
		return &serializer.Response{
			Code:  400,
			Msg:   "创建相册失败",
			Error: fmt.Errorf("user_id 不为当前登录用户").Error(),
		}
	}
	return nil
}

func (svc *Service) AlbumCreate(param *AlbumCreateRequest) serializer.Response {
	userID, ok := svc.ctx.Get("user_id")
	if !ok {
		return serializer.Response{
			Code:  500,
			Msg:   "创建失败",
			Error: fmt.Errorf("当前无登录用户").Error(),
		}
	}
	if err := param.Valid(userID.(uint)); err != nil {
		return *err
	}
	/**
	TODO 验证完 下面是数据库操作 直接调用实体类方法
	*/
	var user model.User
	err := user.GetUserByID(userID.(uint))
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "创建相册失败",
			Error: err.Error(),
		}
	}

	album := model.Album{
		AlbumName: param.AlbumName,
		UserID:    userID.(uint),
	}
	// 调用 album 的 create
	err = album.CreateAlbum()
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "创建相册失败",
			Error: err.Error(),
		}
	}

	// 成功
	return serializer.Response{
		Code: 200,
		Msg:  "创建相册成功",
		Data: serializer.BuildAlbum,
	}
}
