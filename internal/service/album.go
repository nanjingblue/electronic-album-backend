package service

import (
	"electronic-album/internal/dao"
	"electronic-album/internal/model"
	"electronic-album/internal/serializer"
	"fmt"
)

// AlbumCreateService 相册创建服务
type AlbumCreateService struct {
	AlbumName string `form:"album_name" json:"album_name" binding:"required,min=1,max=12"`
	UserID    uint   `form:"user_id" json:"user_id" binding:"required"`
}

func (acr *AlbumCreateService) Valid(userID uint) *serializer.Response {
	if acr.UserID != userID {
		return &serializer.Response{
			Code:  400,
			Msg:   "创建相册失败",
			Error: fmt.Errorf("user_id 不为当前登录用户").Error(),
		}
	}
	return nil
}

func (svc *Service) AlbumCreate(param *AlbumCreateService) serializer.Response {
	//session := sessions.Default(svc.ctx)
	userID := svc.session.Get("user_id")

	if err := param.Valid(userID.(uint)); err != nil {
		return *err
	}

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

	u, _ := svc.ctx.Get("user")

	// 成功
	return serializer.Response{
		Code: 200,
		Msg:  "创建相册成功",
		Data: serializer.BuildAlbum(album, u.(model.User).Username),
	}
}

// AlbumListGetService 获取某个用户相册列表
type AlbumListGetService struct {
	UserID uint `form:"user_id" json:"user_id" binding:"required"`
}

func (algs *AlbumListGetService) Valid() *serializer.Response {
	// 验证 user 是否存在
	_, err := dao.User.GetUserByID(algs.UserID)
	if err != nil {
		return &serializer.Response{
			Code:  400,
			Msg:   "获取相册列表失败",
			Error: err.Error(),
		}
	}
	return nil
}

func (svc *Service) AlbumListGetService(param *AlbumListGetService) serializer.Response {
	err := param.Valid()
	if err != nil {
		return *err
	}
	albums, ok := dao.Album.GetAllAlbumByUserID(param.UserID)
	if ok != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "获取相册列表失败",
			Error: ok.Error(),
		}
	}
	u, _ := dao.User.GetUserByID(param.UserID)
	return serializer.Response{
		Code: 200,
		Msg:  "获取相册列表成功",
		Data: serializer.BuildAlbums(albums, u.Username),
	}
}
