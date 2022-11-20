package v1

import (
	"electronic-album/internal/service"
	"github.com/gin-gonic/gin"
)

func UploadToken(ctx *gin.Context) {
	param := service.UploadTokenService{}
	svc := service.New(ctx)
	if err := ctx.ShouldBind(&param); err == nil {
		res := svc.GetToken(&param)
		ctx.JSON(200, res)
	} else {
		ctx.JSON(400, gin.H{
			"msg":   "获取上传token失败",
			"error": err.Error(),
		})
	}
}
