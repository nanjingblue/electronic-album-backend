package service

import (
	"electronic-gallery/internal/dao"
	"electronic-gallery/internal/model"
	"electronic-gallery/internal/serializer"
	"electronic-gallery/pkg/convert"
	"strings"
)

type PictureService struct{}

type PictureListGetService struct {
	PictureService
	GalleryID uint `form:"gallery_id" json:"gallery_id"`
}

func (pgbgs *PictureListGetService) GetPictures(svc *Service) serializer.Response {
	u, _ := svc.ctx.Get("user")
	user := u.(model.User)

	// 首先判断相册是否存在，并且属于当前用户
	gallery, err := dao.Gallery.GetGalleryByGalleryID(pgbgs.GalleryID)
	if err != nil || gallery.UserID != user.ID {
		return serializer.Response{
			Code:  400,
			Msg:   "相册不存在",
			Error: err.Error(),
		}
	}

	pictures, err := dao.Picture.GetALLPicturesByGalleryID(pgbgs.GalleryID)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "获取图片列表失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildPictures(pictures),
		Msg:  "获取图片列表成功",
	}
}

type PictureCreateService struct {
	PictureService
	GalleryID string `form:"gallery_id" json:"gallery_id" binding:"required"`
	Path      string `form:"path" json:"path" binding:"required"`
}

func (p *PictureCreateService) CreatePicture(svc *Service) serializer.Response {
	u, _ := svc.ctx.Get("user")
	user := u.(model.User)
	galleryID := convert.StrTo(p.GalleryID).MustUInt()
	// 首先判断相册是否存在，并且属于当前用户
	gallery, err := dao.Gallery.GetGalleryByGalleryID(galleryID)
	if err != nil || gallery.UserID != user.ID {
		return serializer.Response{
			Code:  400,
			Msg:   "相册不存在",
			Error: err.Error(),
		}
	}

	filename := p.Path[strings.LastIndex(p.Path, "/")+1:]

	picture := model.Picture{
		PictureName: filename,
		Path:        p.Path,
		UserID:      user.ID,
		GalleryID:   galleryID,
	}

	err = dao.Picture.CreatePicture(&picture)
	if err != nil {
		return serializer.Response{
			Code:  500,
			Msg:   "服务器内部错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 200,
		Data: serializer.BuildPicture(picture),
		Msg:  "添加图片成功",
	}
}
