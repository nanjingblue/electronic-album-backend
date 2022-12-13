package v1

import (
	"electronic-gallery/internal/service"
	"github.com/gin-gonic/gin"
)

// @Summary 获取阿里云上传token
// @Produce json
// @Param file_name body string true "文件名"
// @Param path body string false "路径"
// @Success 200 "成功"
// @Failure 400 "请求错误"
// @Failure 500 "内部错误"
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
