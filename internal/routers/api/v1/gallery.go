package v1

import (
	"electronic-gallery/internal/service"
	"electronic-gallery/pkg/convert"
	"github.com/gin-gonic/gin"
)

// @Summary 相册创建
// @Produce json
// @Param gallery_name body string true "相册名称"
// @Param description body string false "描述"
// @Param cover body string false "封面"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
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

// @Summary 获取相册列表
// @Produce json
// @Param gallery_id body string true "相册id"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
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

// @Summary 相册更新
// @Produce json
// @Param gallery_id body string true "相册id"
// @Param gallery_name body string false "相册名称"
// @Param description body string false "描述"
// @Param cover body string false "封面"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
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

// @Summary 相册删除
// @Produce json
// @Param gallery_id body string true "相册id"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
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
