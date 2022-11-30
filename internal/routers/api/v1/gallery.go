package v1

import (
	"electronic-gallery/internal/service"
	"electronic-gallery/pkg/convert"
	"github.com/gin-gonic/gin"
)

/*
GalleryCreateService 创建相册
*/
func GalleryCreateService(ctx *gin.Context) {
	param := service.GalleryCreateService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.GalleryCreate(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "创建相册失败",
			"error": err.Error(),
		})
	}
}

/*
GalleryGetListService 获取当前用户的所有相册
*/
func GalleryGetListService(ctx *gin.Context) {
	param := service.GalleryListGetService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.GalleryListGetService()
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "获取相册列表失败",
			"error": err.Error(),
		})
	}
}

/*
GalleryUpdateService 更新相册服务
*/
func GalleryUpdateService(ctx *gin.Context) {
	serv := service.GalleryUpdateService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.Update(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "更新相册失败",
			"error": err.Error(),
		})
	}
}

/*
GalleryDeleteService 删除相册服务
*/
func GalleryDeleteService(ctx *gin.Context) {
	serv := service.GalleryDeleteService{
		GalleryID: convert.StrTo(ctx.Param("gallery_id")).MustUInt(),
	}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.Delete(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "删除相册失败",
			"error": err.Error(),
		})
	}
}
