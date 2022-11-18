package v1

import (
	"electronic-album/internal/service"
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

}

/*
GalleryDeleteService 删除相册服务
*/
func GalleryDeleteService(ctx *gin.Context) {

}
