package v1

import (
	"electronic-gallery/internal/service"
	"electronic-gallery/pkg/convert"
	"github.com/gin-gonic/gin"
)

// @Summary 获取某相册图片
// @Produce json
// @Param gallery_id body string true "相册id"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PictureGetPicture(ctx *gin.Context) {
	serv := service.PictureListGetService{
		GalleryID: convert.StrTo(ctx.Param("gallery_id")).MustUInt(),
	}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.GetPictures(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "添加图片失败",
			"error": err.Error(),
		})
	}
}

// @Summary 图片上传
// @Produce json
// @Param gallery_id body string true "相册id"
// @Param path body string true "图片路径"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PictureCreate(ctx *gin.Context) {
	serv := service.PictureCreateService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.CreatePicture(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "添加图片失败",
			"error": err.Error(),
		})
	}
}

// @Summary 图片删除
// @Produce json
// @Param picture_id body string true "图片id"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
func PictureDelete(ctx *gin.Context) {
	serv := service.PictureDeleteService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&serv); err == nil {
		res := serv.Delete(&svc)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "删除图片失败",
			"error": err,
		})
	}
}
