package service

import (
	"electronic-album/internal/dao"
	"electronic-album/internal/model"
	"electronic-album/internal/serializer"
	"electronic-album/pkg/convert"
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
		Cover:       param.Cover,
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
	GalleryID   string `form:"gallery_id" json:"gallery_id" binding:"required"`
	GalleryName string `form:"gallery_name" json:"gallery_name"`
	Description string `form:"description" json:"description"`
	Cover       string `form:"cover" json:"cover"`
}

func (g *GalleryUpdateService) Update(svc *Service) serializer.Response {
	u, _ := svc.ctx.Get("user")
	user := u.(model.User)

	galleryID := convert.StrTo(g.GalleryID).MustUInt()
	gallery, err := dao.Gallery.GetGalleryByGalleryID(galleryID)
	if err != nil {
		return serializer.Response{
			Code:  400,
			Msg:   "更新相册：无相册",
			Error: err.Error(),
		}
	}
	if g.GalleryName != "" {
		gallery.GalleryName = g.GalleryName
	}
	if g.Description != "" {
		gallery.Description = g.GalleryName
	}
	if g.Cover != "" {
		gallery.Cover = g.Cover
	}
	err = dao.Gallery.UpdateGallery(&gallery)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "更新相册：内部错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildGallery(gallery, user.Username),
		Msg:  "更新相册：成功",
	}
}

type GalleryDeleteService struct {
	GalleryID uint `form:"gallery_id" json:"gallery_id" binding:"required"`
}

func (g *GalleryDeleteService) Delete(svc *Service) serializer.Response {
	err := dao.Gallery.DeleteGalleryByGalleryID(g.GalleryID)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "删除相册失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Msg:  "删除相册成功",
	}
}
