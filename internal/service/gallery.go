package service

import (
	"electronic-album/internal/dao"
	"electronic-album/internal/model"
	"electronic-album/internal/serializer"
)

// GalleryCreateService 相册创建服务
type GalleryCreateService struct {
	AlbumName   string `form:"gallery_name" json:"gallery_name" binding:"required,min=1,max=12"`
	Description string `form:"description" json:"description"`
	Cover       string `form:"cover" json:"cover"`
}

func (svc *Service) GalleryCreate(param *GalleryCreateService) serializer.Response {
	u, _ := svc.ctx.Get("user")
	user := u.(model.User)

	gallery := model.Gallery{
		GalleryName: param.AlbumName,
		UserID:      user.ID,
		Cover:       "default_cover.png",
		Status:      model.Active,
		Description: param.Description,
	}

	err := dao.Gallery.CreateGallery(gallery)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Data:  nil,
			Msg:   "创建相册失败",
			Error: err.Error(),
		}
	}

	// 成功
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildGallery(gallery, user.Username),
		Msg:  "创建相册成功",
	}
}

// GalleryListGetService 获取当前用户相册列表
type GalleryListGetService struct{}

func (svc *Service) GalleryListGetService() serializer.Response {
	u, _ := svc.ctx.Get("user")
	user := u.(model.User)

	gallerys, err := dao.Gallery.GetAllGalleryByUserID(user.ID)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "获取相册列表失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "获取相册列表成功",
		Data: serializer.BuildGallerys(gallerys, user.Username),
	}
}

type GalleryUpdateService struct {
	AlbumID uint `form:"album_id" json:"album_id" binding:"required"`
}
