package service

import (
	"electronic-album/internal/dao"
	"electronic-album/internal/model"
	"electronic-album/internal/serializer"
	"fmt"
	"strconv"
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
		Image:     "2ce9c30d-7b7d-469e-a912-4acd72b9d8d2.jpg",
		Status:    1,
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
type AlbumListGetService struct{}

func (svc *Service) AlbumListGetService(id string) serializer.Response {
	uid, _ := strconv.Atoi(id)
	useID := uint(uid)
	albums, ok := dao.Album.GetAllAlbumByUserID(useID)
	if ok != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "获取相册列表失败",
			Error: ok.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "获取相册列表成功",
		Data: serializer.BuildAlbums(albums, "user"),
	}
}

type AlbumUpdateService struct {
	AlbumID uint `form:"album_id" json:"album_id" binding:"required"`
}

//func (svc *Service) AlbumUpdateService(param *AlbumCreateService) serializer.Response {
//r
//}
