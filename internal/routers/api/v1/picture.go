package v1

import (
	"electronic-album/internal/service"
	"electronic-album/pkg/convert"
	"github.com/gin-gonic/gin"
)

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
